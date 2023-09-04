// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: zendesk_service.proto

package zendesk_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ZendeskServiceClient is the client API for ZendeskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZendeskServiceClient interface {
	GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type zendeskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewZendeskServiceClient(cc grpc.ClientConnInterface) ZendeskServiceClient {
	return &zendeskServiceClient{cc}
}

func (c *zendeskServiceClient) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error) {
	out := new(GenerateTokenResponse)
	err := c.cc.Invoke(ctx, "/zendesk_service.ZendeskService/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zendeskServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/zendesk_service.ZendeskService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZendeskServiceServer is the server API for ZendeskService service.
// All implementations must embed UnimplementedZendeskServiceServer
// for forward compatibility
type ZendeskServiceServer interface {
	GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*empty.Empty, error)
	mustEmbedUnimplementedZendeskServiceServer()
}

// UnimplementedZendeskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedZendeskServiceServer struct {
}

func (UnimplementedZendeskServiceServer) GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedZendeskServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedZendeskServiceServer) mustEmbedUnimplementedZendeskServiceServer() {}

// UnsafeZendeskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZendeskServiceServer will
// result in compilation errors.
type UnsafeZendeskServiceServer interface {
	mustEmbedUnimplementedZendeskServiceServer()
}

func RegisterZendeskServiceServer(s grpc.ServiceRegistrar, srv ZendeskServiceServer) {
	s.RegisterService(&ZendeskService_ServiceDesc, srv)
}

func _ZendeskService_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZendeskServiceServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zendesk_service.ZendeskService/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZendeskServiceServer).GenerateToken(ctx, req.(*GenerateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZendeskService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZendeskServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zendesk_service.ZendeskService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZendeskServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ZendeskService_ServiceDesc is the grpc.ServiceDesc for ZendeskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZendeskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zendesk_service.ZendeskService",
	HandlerType: (*ZendeskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateToken",
			Handler:    _ZendeskService_GenerateToken_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _ZendeskService_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zendesk_service.proto",
}
