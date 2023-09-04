package controller

import (
	"awesomeProject4/douyin_grpc/relation_grpc"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	var request relation_grpc.DouYinRelationActionRequest
	request.Token = token
	request.ToUserId, _ = strconv.ParseInt(toUserId, 10, 64)
	newInt, _ := strconv.ParseInt(actionType, 10, 64)
	request.ActionType = int32(newInt)
	conn, err := grpc.Dial(":9004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		return
	}
	defer conn.Close()
	followActionClient := relation_grpc.NewFollowServiceClient(conn)
	resp, err := followActionClient.GiveFollow(context.Background(), &request)

	c.JSON(http.StatusOK, relation_grpc.DouYinRelationActionResponse{
		StatusCode: 0,
		StatusMsg:  resp.StatusMsg,
	})
}
func FollowList(c *gin.Context) {
	var request relation_grpc.DouYinRelationFollowListRequest
	token := c.Query("token")
	userId := c.Query("user_id")
	request.Token = token
	request.UserId, _ = strconv.ParseInt(userId, 10, 64)
	conn, err := grpc.Dial(":9004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		return
	}
	defer conn.Close()
	followListClient := relation_grpc.NewFollowListServiceClient(conn)
	resp, err := followListClient.FollowList(context.Background(), &request)

	c.JSON(http.StatusOK, relation_grpc.DouYinRelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  resp.StatusMsg,
		UserList:   resp.UserList,
	})
}
func FollowerList(c *gin.Context) {
	var request relation_grpc.DouYinRelationFollowerListRequest
	token := c.Query("token")
	userId := c.Query("user_id")
	request.Token = token
	request.UserId, _ = strconv.ParseInt(userId, 10, 64)
	conn, err := grpc.Dial(":9004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		return
	}
	defer conn.Close()
	followerClient := relation_grpc.NewFollowerListServiceClient(conn)
	resp, err := followerClient.FollowerList(context.Background(), &request)

	c.JSON(http.StatusOK, relation_grpc.DouYinRelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  resp.StatusMsg,
		UserList:   resp.UserList,
	})

}
func FriendList(c *gin.Context) {
	var request relation_grpc.DouYinRelationFriendListRequest
	token := c.Query("token")
	userId := c.Query("user_id")
	request.Token = token
	request.UserId, _ = strconv.ParseInt(userId, 10, 64)
	conn, err := grpc.Dial(":9004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		return
	}
	defer conn.Close()
	friendListClient := relation_grpc.NewFriendListServiceClient(conn)
	resp, err := friendListClient.FriendList(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Response": Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		"user_list": resp.UserList,
	})

}
