// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: user.proto

package user_grpc

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
	SignUp_Register_FullMethodName = "/douyin_grpc.SignUp/Register"
)

// SignUpClient is the client API for SignUp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignUpClient interface {
	Register(ctx context.Context, in *DouYinUserRegisterRequest, opts ...grpc.CallOption) (*DouYinUserRegisterResponse, error)
}

type signUpClient struct {
	cc grpc.ClientConnInterface
}

func NewSignUpClient(cc grpc.ClientConnInterface) SignUpClient {
	return &signUpClient{cc}
}

func (c *signUpClient) Register(ctx context.Context, in *DouYinUserRegisterRequest, opts ...grpc.CallOption) (*DouYinUserRegisterResponse, error) {
	out := new(DouYinUserRegisterResponse)
	err := c.cc.Invoke(ctx, SignUp_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignUpServer is the server API for SignUp service.
// All implementations must embed UnimplementedSignUpServer
// for forward compatibility
type SignUpServer interface {
	Register(context.Context, *DouYinUserRegisterRequest) (*DouYinUserRegisterResponse, error)
	mustEmbedUnimplementedSignUpServer()
}

// UnimplementedSignUpServer must be embedded to have forward compatible implementations.
type UnimplementedSignUpServer struct {
}

func (UnimplementedSignUpServer) Register(context.Context, *DouYinUserRegisterRequest) (*DouYinUserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedSignUpServer) mustEmbedUnimplementedSignUpServer() {}

// UnsafeSignUpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignUpServer will
// result in compilation errors.
type UnsafeSignUpServer interface {
	mustEmbedUnimplementedSignUpServer()
}

func RegisterSignUpServer(s grpc.ServiceRegistrar, srv SignUpServer) {
	s.RegisterService(&SignUp_ServiceDesc, srv)
}

func _SignUp_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouYinUserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignUpServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SignUp_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignUpServer).Register(ctx, req.(*DouYinUserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignUp_ServiceDesc is the grpc.ServiceDesc for SignUp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignUp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin_grpc.SignUp",
	HandlerType: (*SignUpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _SignUp_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

const (
	SignIn_Login_FullMethodName = "/douyin_grpc.SignIn/Login"
)

// SignInClient is the client API for SignIn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignInClient interface {
	Login(ctx context.Context, in *DouYinUserLogInRequest, opts ...grpc.CallOption) (*DouYinUserLogInResponse, error)
}

type signInClient struct {
	cc grpc.ClientConnInterface
}

func NewSignInClient(cc grpc.ClientConnInterface) SignInClient {
	return &signInClient{cc}
}

func (c *signInClient) Login(ctx context.Context, in *DouYinUserLogInRequest, opts ...grpc.CallOption) (*DouYinUserLogInResponse, error) {
	out := new(DouYinUserLogInResponse)
	err := c.cc.Invoke(ctx, SignIn_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignInServer is the server API for SignIn service.
// All implementations must embed UnimplementedSignInServer
// for forward compatibility
type SignInServer interface {
	Login(context.Context, *DouYinUserLogInRequest) (*DouYinUserLogInResponse, error)
	mustEmbedUnimplementedSignInServer()
}

// UnimplementedSignInServer must be embedded to have forward compatible implementations.
type UnimplementedSignInServer struct {
}

func (UnimplementedSignInServer) Login(context.Context, *DouYinUserLogInRequest) (*DouYinUserLogInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSignInServer) mustEmbedUnimplementedSignInServer() {}

// UnsafeSignInServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignInServer will
// result in compilation errors.
type UnsafeSignInServer interface {
	mustEmbedUnimplementedSignInServer()
}

func RegisterSignInServer(s grpc.ServiceRegistrar, srv SignInServer) {
	s.RegisterService(&SignIn_ServiceDesc, srv)
}

func _SignIn_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouYinUserLogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignInServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SignIn_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignInServer).Login(ctx, req.(*DouYinUserLogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignIn_ServiceDesc is the grpc.ServiceDesc for SignIn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignIn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin_grpc.SignIn",
	HandlerType: (*SignInServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _SignIn_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

const (
	UserInfo_GetUserInfo_FullMethodName = "/douyin_grpc.UserInfo/GetUserInfo"
)

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoClient interface {
	GetUserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoClient(cc grpc.ClientConnInterface) UserInfoClient {
	return &userInfoClient{cc}
}

func (c *userInfoClient) GetUserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserInfo_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServer is the server API for UserInfo service.
// All implementations must embed UnimplementedUserInfoServer
// for forward compatibility
type UserInfoServer interface {
	GetUserInfo(context.Context, *UserRequest) (*UserResponse, error)
	mustEmbedUnimplementedUserInfoServer()
}

// UnimplementedUserInfoServer must be embedded to have forward compatible implementations.
type UnimplementedUserInfoServer struct {
}

func (UnimplementedUserInfoServer) GetUserInfo(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserInfoServer) mustEmbedUnimplementedUserInfoServer() {}

// UnsafeUserInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoServer will
// result in compilation errors.
type UnsafeUserInfoServer interface {
	mustEmbedUnimplementedUserInfoServer()
}

func RegisterUserInfoServer(s grpc.ServiceRegistrar, srv UserInfoServer) {
	s.RegisterService(&UserInfo_ServiceDesc, srv)
}

func _UserInfo_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfo_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).GetUserInfo(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfo_ServiceDesc is the grpc.ServiceDesc for UserInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin_grpc.UserInfo",
	HandlerType: (*UserInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _UserInfo_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
