/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package recordingmerger

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	profilebase "sigs.k8s.io/security-profiles-operator/api/profilebase/v1alpha1"
	profilerecording1alpha1 "sigs.k8s.io/security-profiles-operator/api/profilerecording/v1alpha1"
	seccompprofile "sigs.k8s.io/security-profiles-operator/api/seccompprofile/v1beta1"
	selinuxprofileapi "sigs.k8s.io/security-profiles-operator/api/selinuxprofile/v1alpha2"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/util"
)

func mergedObjectMeta(profileName, recordingName, namespace string) *metav1.ObjectMeta {
	return &metav1.ObjectMeta{
		Name:      profileName,
		Namespace: namespace,
		Labels: map[string]string{
			profilerecording1alpha1.ProfileToRecordingLabel: recordingName,
		},
	}
}

func mergedProfileName(recordingName string, prf metav1.Object) string {
	suffix := prf.GetLabels()[profilerecording1alpha1.ProfileToContainerLabel]
	if suffix == "" {
		suffix = prf.GetName()
	}
	return fmt.Sprintf("%s-%s", recordingName, suffix)
}

func mergeProfiles(profiles []mergeableProfile) (mergeableProfile, error) {
	if len(profiles) == 0 {
		return nil, fmt.Errorf("cannot merge empty list of profiles")
	}

	base := profiles[0]
	if len(profiles) == 1 {
		return base, nil
	}

	mergeSlice := profiles[1:]
	for i := range mergeSlice {
		err := base.merge(mergeSlice[i])
		if err != nil {
			return nil, fmt.Errorf("failed to merge profile %s: %w", mergeSlice[i].GetName(), err)
		}
	}

	return base, nil
}

func listPartialProfiles(
	ctx context.Context,
	cli client.Client,
	list client.ObjectList,
	recording *profilerecording1alpha1.ProfileRecording,
) ([]mergeableProfile, error) {
	if err := cli.List(
		ctx,
		list,
		client.InNamespace(recording.Namespace),
		client.MatchingLabels{
			profilerecording1alpha1.ProfileToRecordingLabel: recording.Name,
			profilebase.ProfilePartialLabel:                 "true",
		}); err != nil {
		return nil, fmt.Errorf("listing partial profiles for %s: %w", recording.Name, err)
	}

	partialProfiles := make([]mergeableProfile, 0)
	if err := meta.EachListItem(list, func(obj runtime.Object) error {
		clientObj, ok := obj.(client.Object)
		if !ok {
			return fmt.Errorf("object %T is not a client.Object", obj)
		}

		partialPrf, err := newMergeableProfile(clientObj)
		if err != nil {
			return fmt.Errorf("failed to create mergeable profile for %s: %w", clientObj.GetName(), err)
		}

		partialProfiles = append(partialProfiles, partialPrf)
		return nil
	}); err != nil {
		return nil, fmt.Errorf("iterating over partial profiles: %w", err)
	}

	return partialProfiles, nil
}

func deletePartialProfiles(
	ctx context.Context,
	cli client.Client,
	prf client.Object,
	recording *profilerecording1alpha1.ProfileRecording,
) error {
	return cli.DeleteAllOf(
		ctx,
		prf,
		client.InNamespace(recording.Namespace),
		client.MatchingLabels{
			profilerecording1alpha1.ProfileToRecordingLabel: recording.Name,
			profilebase.ProfilePartialLabel:                 "true",
		})
}

func newMergeableProfile(obj client.Object) (mergeableProfile, error) {
	switch obj := obj.(type) {
	case *seccompprofile.SeccompProfile:
		return &mergeableSeccompProfile{SeccompProfile: *obj}, nil
	case *selinuxprofileapi.SelinuxProfile:
		return &MergeableSelinuxProfile{SelinuxProfile: *obj}, nil
	default:
		return nil, fmt.Errorf("cannot convert %T to mergeableProfile", obj)
	}
}

type mergeableProfile interface {
	client.Object

	merge(profile mergeableProfile) error
	getProfile() client.Object
}

type mergeableSeccompProfile struct {
	seccompprofile.SeccompProfile
}

func (sp *mergeableSeccompProfile) merge(other mergeableProfile) error {
	otherSP, ok := other.(*mergeableSeccompProfile)
	if !ok {
		return fmt.Errorf("cannot merge SeccompProfile with %T", other)
	}
	sp.Spec.Syscalls = util.UnionSyscalls(sp.Spec.Syscalls, otherSP.Spec.Syscalls)

	return nil
}

func (sp *mergeableSeccompProfile) getProfile() client.Object {
	return &sp.SeccompProfile
}

type MergeableSelinuxProfile struct {
	selinuxprofileapi.SelinuxProfile
}

func (sp *MergeableSelinuxProfile) getProfile() client.Object {
	return &sp.SelinuxProfile
}

func (sp *MergeableSelinuxProfile) merge(other mergeableProfile) error {
	// TODO(jhrozek): should we be defensive about checking if other attributes match as well? (e.g. inherit)
	otherSP, ok := other.(*MergeableSelinuxProfile)
	if !ok {
		return fmt.Errorf("cannot merge selinuxProfile with %T", other)
	}
	sp.Spec.Allow = addAllow(sp.Spec.Allow, otherSP.Spec.Allow)

	return nil
}

func addAllow(union, additional selinuxprofileapi.Allow) selinuxprofileapi.Allow {
	for labelKey, permMap := range additional {
		if _, ok := union[labelKey]; !ok {
			union[labelKey] = make(map[selinuxprofileapi.ObjectClassKey]selinuxprofileapi.PermissionSet)
		}
		for objClass, perms := range permMap {
			union[labelKey][objClass] = append(union[labelKey][objClass], perms...)
		}
	}

	return union
}
