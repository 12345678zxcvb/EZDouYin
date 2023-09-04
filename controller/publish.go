package controller

import (
	"awesomeProject4/Logic"
	"awesomeProject4/douyin_grpc/publish_grpc"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"path/filepath"
	"strconv"
)

func Publish(c *gin.Context) {
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	token := c.PostForm("token")
	if err != nil {
		msg := err.Error()
		c.JSON(http.StatusOK, publish_grpc.DouYinPublishActionResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
		})
		return
	}
	grpcConn, err := grpc.Dial(":9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		return
	}
	defer grpcConn.Close()
	publishActionClient := publish_grpc.NewActionPublishClient(grpcConn)
	resp, err := publishActionClient.PublishAction(context.Background(), &publish_grpc.DouYinPublishActionRequest{})
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("_%s", filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to save file"})
		return
	}
	secretKey := []byte("zhang")
	username, err := Logic.ExtractUsernameFromToken(token, secretKey)
	Logic.PublishVideo(username, saveFile, title, finalName)
	c.JSON(http.StatusOK, publish_grpc.DouYinPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  resp.StatusMsg,
	})

}
func PublishList(c *gin.Context) {
	id := c.Query("user_id")
	newId, _ := strconv.ParseInt(id, 10, 64)
	var request publish_grpc.DouYinPublishListRequest
	request.UserId, _ = strconv.ParseInt(id, 10, 64)
	grpcConn, err := grpc.Dial(":9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		return
	}
	defer grpcConn.Close()
	PublishListClient := publish_grpc.NewListPublishClient(grpcConn)
	req := publish_grpc.DouYinPublishListRequest{
		UserId: newId,
	}
	resp, err := PublishListClient.PublishList(context.Background(), &req)

	c.JSON(http.StatusOK, publish_grpc.DouYinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  resp.StatusMsg,
		VideoList:  resp.VideoList,
	})
}
