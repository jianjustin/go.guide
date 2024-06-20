// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0--rc1
// source: a/a.proto

package proto

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

const (
	A_AddAPreffix_FullMethodName = "/proto.A/AddAPreffix"
)

// AClient is the client API for A service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AClient interface {
	// Sends a greeting
	AddAPreffix(ctx context.Context, in *ARequest, opts ...grpc.CallOption) (*AReply, error)
}

type aClient struct {
	cc grpc.ClientConnInterface
}

func NewAClient(cc grpc.ClientConnInterface) AClient {
	return &aClient{cc}
}

func (c *aClient) AddAPreffix(ctx context.Context, in *ARequest, opts ...grpc.CallOption) (*AReply, error) {
	out := new(AReply)
	err := c.cc.Invoke(ctx, A_AddAPreffix_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AServer is the server API for A service.
// All implementations should embed UnimplementedAServer
// for forward compatibility
type AServer interface {
	// Sends a greeting
	AddAPreffix(context.Context, *ARequest) (*AReply, error)
}

// UnimplementedAServer should be embedded to have forward compatible implementations.
type UnimplementedAServer struct {
}

func (UnimplementedAServer) AddAPreffix(context.Context, *ARequest) (*AReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAPreffix not implemented")
}

// UnsafeAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AServer will
// result in compilation errors.
type UnsafeAServer interface {
	mustEmbedUnimplementedAServer()
}

func RegisterAServer(s grpc.ServiceRegistrar, srv AServer) {
	s.RegisterService(&A_ServiceDesc, srv)
}

func _A_AddAPreffix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AServer).AddAPreffix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: A_AddAPreffix_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AServer).AddAPreffix(ctx, req.(*ARequest))
	}
	return interceptor(ctx, in, info, handler)
}

// A_ServiceDesc is the grpc.ServiceDesc for A service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var A_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.A",
	HandlerType: (*AServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAPreffix",
			Handler:    _A_AddAPreffix_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "a/a.proto",
}
