package user_grpc

import (
	"awesomeProject4/Logic"
	"awesomeProject4/common"
	"awesomeProject4/common/model"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GetUserInfo struct {
}
type SignInService struct {
}
type SignUpService struct{}

func (i *GetUserInfo) mustEmbedUnimplementedUserInfoServer() {
	//TODO implement me
	panic("implement me")
}

func (*GetUserInfo) GetUserInfo(ctx context.Context, req *UserRequest) (*UserResponse, error) {

	//var user

	return &UserResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

func (SignInService) mustEmbedUnimplementedSignInServer() {
	//TODO implement me
	panic("implement me")
}

func (SignInService) Login(ctx context.Context, req *DouYinUserLogInRequest) (*DouYinUserLogInResponse, error) {
	db := common.GetDB()
	var userInfoTable model.UserInfoTable
	var value string
	value = ""
	db.Take(&userInfoTable, "Name = ?", req.Username)
	if req.Password != userInfoTable.PassWord {
		return &DouYinUserLogInResponse{
			StatusCode: 0,
			StatusMsg:  &value,
			UserId:     userInfoTable.ID,
			Token:      "",
		}, nil
	}
	TokenString, err := Logic.GenerateToken(req.Username, req.Password)
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	value = "success"
	return &DouYinUserLogInResponse{
		StatusCode: 0,
		StatusMsg:  &value,
		UserId:     userInfoTable.ID,
		Token:      TokenString,
	}, nil
}

func (SignUpService) mustEmbedUnimplementedSignUpServer() {
	panic("implement me")
}
func (SignUpService) Register(ctx context.Context, req *DouYinUserRegisterRequest) (*DouYinUserRegisterResponse, error) {
	TokenString, err := Logic.GenerateToken(req.Username, req.Password)
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	value := "success"
	db := common.GetDB()
	var userInfoTable model.UserInfoTable
	db.Take(&userInfoTable, "Name = ?", req.Username)
	return &DouYinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  &value,
		UserId:     userInfoTable.ID,
		Token:      TokenString,
	}, nil
}

func StartGRPCServer() {
	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()

	// 注册 SignUpService
	signUpService := &SignUpService{}
	signInService := &SignInService{}
	getUserInfo := &GetUserInfo{}
	RegisterSignUpServer(grpcServer, signUpService)
	RegisterSignInServer(grpcServer, signInService)
	RegisterUserInfoServer(grpcServer, getUserInfo)

	// 启动 gRPC 服务器
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
