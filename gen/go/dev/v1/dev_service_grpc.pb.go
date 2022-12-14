// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: dev/v1/dev_service.proto

package devv1

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

// DevServiceClient is the client API for DevService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevServiceClient interface {
	NewProject(ctx context.Context, in *NewProjectRequest, opts ...grpc.CallOption) (*NewProjectResponse, error)
}

type devServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDevServiceClient(cc grpc.ClientConnInterface) DevServiceClient {
	return &devServiceClient{cc}
}

func (c *devServiceClient) NewProject(ctx context.Context, in *NewProjectRequest, opts ...grpc.CallOption) (*NewProjectResponse, error) {
	out := new(NewProjectResponse)
	err := c.cc.Invoke(ctx, "/dev.v1.DevService/NewProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevServiceServer is the server API for DevService service.
// All implementations must embed UnimplementedDevServiceServer
// for forward compatibility
type DevServiceServer interface {
	NewProject(context.Context, *NewProjectRequest) (*NewProjectResponse, error)
	mustEmbedUnimplementedDevServiceServer()
}

// UnimplementedDevServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDevServiceServer struct {
}

func (UnimplementedDevServiceServer) NewProject(context.Context, *NewProjectRequest) (*NewProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewProject not implemented")
}
func (UnimplementedDevServiceServer) mustEmbedUnimplementedDevServiceServer() {}

// UnsafeDevServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevServiceServer will
// result in compilation errors.
type UnsafeDevServiceServer interface {
	mustEmbedUnimplementedDevServiceServer()
}

func RegisterDevServiceServer(s grpc.ServiceRegistrar, srv DevServiceServer) {
	s.RegisterService(&DevService_ServiceDesc, srv)
}

func _DevService_NewProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevServiceServer).NewProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.v1.DevService/NewProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevServiceServer).NewProject(ctx, req.(*NewProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DevService_ServiceDesc is the grpc.ServiceDesc for DevService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DevService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.v1.DevService",
	HandlerType: (*DevServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewProject",
			Handler:    _DevService_NewProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dev/v1/dev_service.proto",
}
