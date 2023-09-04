package controller

import (
	"awesomeProject4/douyin_grpc/favorite_grpc"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	var request favorite_grpc.DouYinFavoriteActionRequest
	request.Token = token
	request.VideoId, _ = strconv.ParseInt(videoId, 10, 64)
	newInt, _ := strconv.ParseInt(actionType, 10, 64)
	request.ActionType = int32(newInt)
	conn, err := grpc.Dial(":9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to connect grpc",
		})
	}
	defer conn.Close()
	favoriteClient := favorite_grpc.NewFavoriteActionClient(conn)
	resp, err := favoriteClient.GiveFavoriteService(context.Background(), &request)
	c.JSON(http.StatusOK, &favorite_grpc.DouYinFavoriteActionResponse{
		StatusCode: 0,
		StatusMsg:  resp.StatusMsg,
	})
}
func FavoriteList(c *gin.Context) {

	var request favorite_grpc.DouYinFavoriteListRequest
	token := c.Query("Token")
	userId := c.Query("user_id")
	request.UserId, _ = strconv.ParseInt(userId, 10, 64)
	request.Token = token
	conn, err := grpc.Dial(":9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to connect grpc",
		})
	}
	defer conn.Close()
	favoriteClient := favorite_grpc.NewFavoriteListClient(conn)
	resp, err := favoriteClient.FavoriteListService(context.Background(), &request)
	c.JSON(http.StatusOK, &favorite_grpc.DouYinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  resp.StatusMsg,
		VideoList:  resp.VideoList,
	})
}
