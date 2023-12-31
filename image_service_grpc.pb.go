// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: image_service.proto

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

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	GetSingleImg(ctx context.Context, in *ResizeRequest, opts ...grpc.CallOption) (*GetSingleImgResult, error)
	GetMultiImgs(ctx context.Context, in *Imgrequests, opts ...grpc.CallOption) (ImageService_GetMultiImgsClient, error)
	Getserverstatus(ctx context.Context, in *GetServerStatusRequest, opts ...grpc.CallOption) (*Serverstatus, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) GetSingleImg(ctx context.Context, in *ResizeRequest, opts ...grpc.CallOption) (*GetSingleImgResult, error) {
	out := new(GetSingleImgResult)
	err := c.cc.Invoke(ctx, "/image_service.ImageService/GetSingleImg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) GetMultiImgs(ctx context.Context, in *Imgrequests, opts ...grpc.CallOption) (ImageService_GetMultiImgsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageService_ServiceDesc.Streams[0], "/image_service.ImageService/getMultiImgs", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageServiceGetMultiImgsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ImageService_GetMultiImgsClient interface {
	Recv() (*Imgresult, error)
	grpc.ClientStream
}

type imageServiceGetMultiImgsClient struct {
	grpc.ClientStream
}

func (x *imageServiceGetMultiImgsClient) Recv() (*Imgresult, error) {
	m := new(Imgresult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imageServiceClient) Getserverstatus(ctx context.Context, in *GetServerStatusRequest, opts ...grpc.CallOption) (*Serverstatus, error) {
	out := new(Serverstatus)
	err := c.cc.Invoke(ctx, "/image_service.ImageService/Getserverstatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility
type ImageServiceServer interface {
	GetSingleImg(context.Context, *ResizeRequest) (*GetSingleImgResult, error)
	GetMultiImgs(*Imgrequests, ImageService_GetMultiImgsServer) error
	Getserverstatus(context.Context, *GetServerStatusRequest) (*Serverstatus, error)
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (UnimplementedImageServiceServer) GetSingleImg(context.Context, *ResizeRequest) (*GetSingleImgResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleImg not implemented")
}
func (UnimplementedImageServiceServer) GetMultiImgs(*Imgrequests, ImageService_GetMultiImgsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMultiImgs not implemented")
}
func (UnimplementedImageServiceServer) Getserverstatus(context.Context, *GetServerStatusRequest) (*Serverstatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getserverstatus not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_GetSingleImg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).GetSingleImg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image_service.ImageService/GetSingleImg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).GetSingleImg(ctx, req.(*ResizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_GetMultiImgs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Imgrequests)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImageServiceServer).GetMultiImgs(m, &imageServiceGetMultiImgsServer{stream})
}

type ImageService_GetMultiImgsServer interface {
	Send(*Imgresult) error
	grpc.ServerStream
}

type imageServiceGetMultiImgsServer struct {
	grpc.ServerStream
}

func (x *imageServiceGetMultiImgsServer) Send(m *Imgresult) error {
	return x.ServerStream.SendMsg(m)
}

func _ImageService_Getserverstatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Getserverstatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image_service.ImageService/Getserverstatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Getserverstatus(ctx, req.(*GetServerStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "image_service.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSingleImg",
			Handler:    _ImageService_GetSingleImg_Handler,
		},
		{
			MethodName: "Getserverstatus",
			Handler:    _ImageService_Getserverstatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getMultiImgs",
			Handler:       _ImageService_GetMultiImgs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "image_service.proto",
}
