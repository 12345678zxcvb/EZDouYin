package controller

import (
	"awesomeProject4/Logic"
	"awesomeProject4/common"
	"awesomeProject4/common/model"
	"awesomeProject4/douyin_grpc/user_grpc"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func Register(c *gin.Context) {
	var request user_grpc.DouYinUserRegisterRequest
	username := c.Query("username")
	password := c.Query("password")
	request.Username = username
	request.Password = password
	str := Logic.RegisterVerification(username, password)

	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connection to GRPC server"})
	}
	defer conn.Close()
	signUpClient := user_grpc.NewSignUpClient(conn)
	response, err := signUpClient.Register(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
	}
	message := &str
	if str != "" {
		c.JSON(http.StatusOK, user_grpc.DouYinUserLogInResponse{
			StatusCode: 1,
			StatusMsg:  message,
		})
	} else {
		c.JSON(http.StatusOK, user_grpc.DouYinUserRegisterResponse{
			StatusCode: 0,
			StatusMsg:  response.StatusMsg,
			UserId:     response.UserId,
			Token:      response.Token,
		})
	}
}

func Login(c *gin.Context) {
	var request user_grpc.DouYinUserLogInRequest
	username := c.Query("username")
	password := c.Query("password")
	request.Username = username
	request.Password = password
	// 调用 gRPC 服务处理登录逻辑
	grpcConn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		return
	}
	defer grpcConn.Close()
	signInClient := user_grpc.NewSignInClient(grpcConn)
	response, err := signInClient.Login(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}
	str := Logic.LogInVerification(username, password)
	if str != "" {
		c.JSON(http.StatusOK, user_grpc.DouYinUserLogInResponse{
			StatusCode: 1,
			StatusMsg:  &str,
		})
	} else {
		c.JSON(http.StatusOK, user_grpc.DouYinUserLogInResponse{
			StatusCode: 0,
			StatusMsg:  response.StatusMsg,
			UserId:     response.UserId,
			Token:      response.Token,
		})
	}
}

func UserInfo(c *gin.Context) {
	id := c.Query("user_id")
	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		return
	}
	var userInfoRequest user_grpc.UserRequest
	defer conn.Close()
	userInfoClient := user_grpc.NewUserInfoClient(conn)
	response, err := userInfoClient.GetUserInfo(context.Background(), &userInfoRequest)
	var userInfoTable model.UserInfoTable
	db := common.GetDB()
	db.Take(&userInfoTable, id)
	//response.StatusMsg = "success"
	user := ConvertToUser(userInfoTable)
	c.JSON(http.StatusOK, user_grpc.UserResponse{
		StatusCode: 0,
		StatusMsg:  response.StatusMsg,
		User:       &user,
	})

}
func ConvertToUser(userinfotable model.UserInfoTable) user_grpc.User {
	user := user_grpc.User{
		Id:              userinfotable.ID,
		Name:            userinfotable.Name,
		FollowCount:     userinfotable.FollowCount,
		FollowerCount:   userinfotable.FollowerCount,
		IsFollow:        userinfotable.IsFollow,
		Avatar:          userinfotable.Avatar,
		BackgroundImage: userinfotable.BackgroundImage,
		Signature:       userinfotable.Signature,
		TotalFavorited:  userinfotable.TotalFavorite,
		WorkCount:       userinfotable.WorkCount,
		FavoriteCount:   userinfotable.FavoriteCount,
	}
	return user
}
