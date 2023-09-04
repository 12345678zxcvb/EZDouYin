// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: comment.proto

package comment_grpc

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
	CommentService_GiveComment_FullMethodName = "/comment_grpc.CommentService/GiveComment"
)

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	GiveComment(ctx context.Context, in *DouYinCommentActionRequest, opts ...grpc.CallOption) (*DouYinCommentActionResponse, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) GiveComment(ctx context.Context, in *DouYinCommentActionRequest, opts ...grpc.CallOption) (*DouYinCommentActionResponse, error) {
	out := new(DouYinCommentActionResponse)
	err := c.cc.Invoke(ctx, CommentService_GiveComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility
type CommentServiceServer interface {
	GiveComment(context.Context, *DouYinCommentActionRequest) (*DouYinCommentActionResponse, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (UnimplementedCommentServiceServer) GiveComment(context.Context, *DouYinCommentActionRequest) (*DouYinCommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GiveComment not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_GiveComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouYinCommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GiveComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GiveComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GiveComment(ctx, req.(*DouYinCommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment_grpc.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GiveComment",
			Handler:    _CommentService_GiveComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}

const (
	CommentListService_ListComment_FullMethodName = "/comment_grpc.CommentListService/ListComment"
)

// CommentListServiceClient is the client API for CommentListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentListServiceClient interface {
	ListComment(ctx context.Context, in *DouYinCommentListRequest, opts ...grpc.CallOption) (*DouYinCommentListResponse, error)
}

type commentListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentListServiceClient(cc grpc.ClientConnInterface) CommentListServiceClient {
	return &commentListServiceClient{cc}
}

func (c *commentListServiceClient) ListComment(ctx context.Context, in *DouYinCommentListRequest, opts ...grpc.CallOption) (*DouYinCommentListResponse, error) {
	out := new(DouYinCommentListResponse)
	err := c.cc.Invoke(ctx, CommentListService_ListComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentListServiceServer is the server API for CommentListService service.
// All implementations must embed UnimplementedCommentListServiceServer
// for forward compatibility
type CommentListServiceServer interface {
	ListComment(context.Context, *DouYinCommentListRequest) (*DouYinCommentListResponse, error)
	mustEmbedUnimplementedCommentListServiceServer()
}

// UnimplementedCommentListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentListServiceServer struct {
}

func (UnimplementedCommentListServiceServer) ListComment(context.Context, *DouYinCommentListRequest) (*DouYinCommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComment not implemented")
}
func (UnimplementedCommentListServiceServer) mustEmbedUnimplementedCommentListServiceServer() {}

// UnsafeCommentListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentListServiceServer will
// result in compilation errors.
type UnsafeCommentListServiceServer interface {
	mustEmbedUnimplementedCommentListServiceServer()
}

func RegisterCommentListServiceServer(s grpc.ServiceRegistrar, srv CommentListServiceServer) {
	s.RegisterService(&CommentListService_ServiceDesc, srv)
}

func _CommentListService_ListComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouYinCommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentListServiceServer).ListComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentListService_ListComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentListServiceServer).ListComment(ctx, req.(*DouYinCommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentListService_ServiceDesc is the grpc.ServiceDesc for CommentListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment_grpc.CommentListService",
	HandlerType: (*CommentListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListComment",
			Handler:    _CommentListService_ListComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
