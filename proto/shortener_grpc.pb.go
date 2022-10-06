// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: shortener.proto

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

// UrlsClient is the client API for Urls service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlsClient interface {
	CreateShortURLHandler(ctx context.Context, in *AddUrlRequest, opts ...grpc.CallOption) (*AddUrlResponse, error)
	GetShortURLHandler(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error)
}

type urlsClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlsClient(cc grpc.ClientConnInterface) UrlsClient {
	return &urlsClient{cc}
}

func (c *urlsClient) CreateShortURLHandler(ctx context.Context, in *AddUrlRequest, opts ...grpc.CallOption) (*AddUrlResponse, error) {
	out := new(AddUrlResponse)
	err := c.cc.Invoke(ctx, "/shortener.Urls/CreateShortURLHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlsClient) GetShortURLHandler(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error) {
	out := new(GetUrlResponse)
	err := c.cc.Invoke(ctx, "/shortener.Urls/GetShortURLHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlsServer is the server API for Urls service.
// All implementations must embed UnimplementedUrlsServer
// for forward compatibility
type UrlsServer interface {
	CreateShortURLHandler(context.Context, *AddUrlRequest) (*AddUrlResponse, error)
	GetShortURLHandler(context.Context, *GetUrlRequest) (*GetUrlResponse, error)
	mustEmbedUnimplementedUrlsServer()
}

// UnimplementedUrlsServer must be embedded to have forward compatible implementations.
type UnimplementedUrlsServer struct {
}

func (UnimplementedUrlsServer) CreateShortURLHandler(context.Context, *AddUrlRequest) (*AddUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShortURLHandler not implemented")
}
func (UnimplementedUrlsServer) GetShortURLHandler(context.Context, *GetUrlRequest) (*GetUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortURLHandler not implemented")
}
func (UnimplementedUrlsServer) mustEmbedUnimplementedUrlsServer() {}

// UnsafeUrlsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlsServer will
// result in compilation errors.
type UnsafeUrlsServer interface {
	mustEmbedUnimplementedUrlsServer()
}

func RegisterUrlsServer(s grpc.ServiceRegistrar, srv UrlsServer) {
	s.RegisterService(&Urls_ServiceDesc, srv)
}

func _Urls_CreateShortURLHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlsServer).CreateShortURLHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Urls/CreateShortURLHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlsServer).CreateShortURLHandler(ctx, req.(*AddUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Urls_GetShortURLHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlsServer).GetShortURLHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Urls/GetShortURLHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlsServer).GetShortURLHandler(ctx, req.(*GetUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Urls_ServiceDesc is the grpc.ServiceDesc for Urls service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Urls_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortener.Urls",
	HandlerType: (*UrlsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShortURLHandler",
			Handler:    _Urls_CreateShortURLHandler_Handler,
		},
		{
			MethodName: "GetShortURLHandler",
			Handler:    _Urls_GetShortURLHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shortener.proto",
}
