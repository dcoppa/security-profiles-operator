//
//Copyright 2021 The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: api/grpc/enricher/api.proto

package api_enricher

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Enricher_Syscalls_FullMethodName      = "/api_enricher.Enricher/Syscalls"
	Enricher_ResetSyscalls_FullMethodName = "/api_enricher.Enricher/ResetSyscalls"
	Enricher_Avcs_FullMethodName          = "/api_enricher.Enricher/Avcs"
	Enricher_ResetAvcs_FullMethodName     = "/api_enricher.Enricher/ResetAvcs"
)

// EnricherClient is the client API for Enricher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnricherClient interface {
	Syscalls(ctx context.Context, in *SyscallsRequest, opts ...grpc.CallOption) (*SyscallsResponse, error)
	ResetSyscalls(ctx context.Context, in *SyscallsRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	Avcs(ctx context.Context, in *AvcRequest, opts ...grpc.CallOption) (*AvcResponse, error)
	ResetAvcs(ctx context.Context, in *AvcRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type enricherClient struct {
	cc grpc.ClientConnInterface
}

func NewEnricherClient(cc grpc.ClientConnInterface) EnricherClient {
	return &enricherClient{cc}
}

func (c *enricherClient) Syscalls(ctx context.Context, in *SyscallsRequest, opts ...grpc.CallOption) (*SyscallsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyscallsResponse)
	err := c.cc.Invoke(ctx, Enricher_Syscalls_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enricherClient) ResetSyscalls(ctx context.Context, in *SyscallsRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Enricher_ResetSyscalls_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enricherClient) Avcs(ctx context.Context, in *AvcRequest, opts ...grpc.CallOption) (*AvcResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AvcResponse)
	err := c.cc.Invoke(ctx, Enricher_Avcs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enricherClient) ResetAvcs(ctx context.Context, in *AvcRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Enricher_ResetAvcs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnricherServer is the server API for Enricher service.
// All implementations must embed UnimplementedEnricherServer
// for forward compatibility.
type EnricherServer interface {
	Syscalls(context.Context, *SyscallsRequest) (*SyscallsResponse, error)
	ResetSyscalls(context.Context, *SyscallsRequest) (*EmptyResponse, error)
	Avcs(context.Context, *AvcRequest) (*AvcResponse, error)
	ResetAvcs(context.Context, *AvcRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedEnricherServer()
}

// UnimplementedEnricherServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEnricherServer struct{}

func (UnimplementedEnricherServer) Syscalls(context.Context, *SyscallsRequest) (*SyscallsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Syscalls not implemented")
}
func (UnimplementedEnricherServer) ResetSyscalls(context.Context, *SyscallsRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetSyscalls not implemented")
}
func (UnimplementedEnricherServer) Avcs(context.Context, *AvcRequest) (*AvcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Avcs not implemented")
}
func (UnimplementedEnricherServer) ResetAvcs(context.Context, *AvcRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetAvcs not implemented")
}
func (UnimplementedEnricherServer) mustEmbedUnimplementedEnricherServer() {}
func (UnimplementedEnricherServer) testEmbeddedByValue()                  {}

// UnsafeEnricherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnricherServer will
// result in compilation errors.
type UnsafeEnricherServer interface {
	mustEmbedUnimplementedEnricherServer()
}

func RegisterEnricherServer(s grpc.ServiceRegistrar, srv EnricherServer) {
	// If the following call pancis, it indicates UnimplementedEnricherServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Enricher_ServiceDesc, srv)
}

func _Enricher_Syscalls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyscallsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnricherServer).Syscalls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Enricher_Syscalls_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnricherServer).Syscalls(ctx, req.(*SyscallsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enricher_ResetSyscalls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyscallsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnricherServer).ResetSyscalls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Enricher_ResetSyscalls_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnricherServer).ResetSyscalls(ctx, req.(*SyscallsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enricher_Avcs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnricherServer).Avcs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Enricher_Avcs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnricherServer).Avcs(ctx, req.(*AvcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enricher_ResetAvcs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnricherServer).ResetAvcs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Enricher_ResetAvcs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnricherServer).ResetAvcs(ctx, req.(*AvcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Enricher_ServiceDesc is the grpc.ServiceDesc for Enricher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Enricher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api_enricher.Enricher",
	HandlerType: (*EnricherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Syscalls",
			Handler:    _Enricher_Syscalls_Handler,
		},
		{
			MethodName: "ResetSyscalls",
			Handler:    _Enricher_ResetSyscalls_Handler,
		},
		{
			MethodName: "Avcs",
			Handler:    _Enricher_Avcs_Handler,
		},
		{
			MethodName: "ResetAvcs",
			Handler:    _Enricher_ResetAvcs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/enricher/api.proto",
}
