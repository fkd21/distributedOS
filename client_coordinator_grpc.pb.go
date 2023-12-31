// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: client_coordinator.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoordinatorClient interface {
	GetBestServer(ctx context.Context, in *GetBestServerRequest, opts ...grpc.CallOption) (*GetBestServerResponse, error)
}

type coordinatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinatorClient(cc grpc.ClientConnInterface) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) GetBestServer(ctx context.Context, in *GetBestServerRequest, opts ...grpc.CallOption) (*GetBestServerResponse, error) {
	out := new(GetBestServerResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/GetBestServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
// All implementations must embed UnimplementedCoordinatorServer
// for forward compatibility
type CoordinatorServer interface {
	GetBestServer(context.Context, *GetBestServerRequest) (*GetBestServerResponse, error)
	mustEmbedUnimplementedCoordinatorServer()
}

// UnimplementedCoordinatorServer must be embedded to have forward compatible implementations.
type UnimplementedCoordinatorServer struct {
}

func (UnimplementedCoordinatorServer) GetBestServer(context.Context, *GetBestServerRequest) (*GetBestServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBestServer not implemented")
}
func (UnimplementedCoordinatorServer) mustEmbedUnimplementedCoordinatorServer() {}

// UnsafeCoordinatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoordinatorServer will
// result in compilation errors.
type UnsafeCoordinatorServer interface {
	mustEmbedUnimplementedCoordinatorServer()
}

func RegisterCoordinatorServer(s grpc.ServiceRegistrar, srv CoordinatorServer) {
	s.RegisterService(&Coordinator_ServiceDesc, srv)
}

func _Coordinator_GetBestServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBestServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).GetBestServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coordinator.Coordinator/GetBestServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).GetBestServer(ctx, req.(*GetBestServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Coordinator_ServiceDesc is the grpc.ServiceDesc for Coordinator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Coordinator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coordinator.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBestServer",
			Handler:    _Coordinator_GetBestServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_coordinator.proto",
}
