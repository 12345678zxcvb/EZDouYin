package controller

import (
	"awesomeProject4/douyin_grpc/comment_grpc"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	var request comment_grpc.DouYinCommentActionRequest
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentId := c.Query("comment_id")
	request.Token = token
	request.VideoId, _ = strconv.ParseInt(videoId, 10, 64)
	newInt, err := strconv.ParseInt(actionType, 10, 64)
	request.ActionType = int32(newInt)
	request.CommentText = &commentText
	newInt1, _ := strconv.ParseInt(commentId, 10, 64)
	request.CommentId = &newInt1
	conn, err := grpc.Dial(":9003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to connect",
		})
	}
	defer conn.Close()
	commentClient := comment_grpc.NewCommentServiceClient(conn)
	resp, err := commentClient.GiveComment(context.Background(), &request)
	c.JSON(http.StatusOK, &comment_grpc.DouYinCommentActionResponse{
		StatusCode: 0,
		StatusMsg:  resp.StatusMsg,
		Comment:    resp.Comment,
	})
}
func CommentList(c *gin.Context) {
	var request comment_grpc.DouYinCommentListRequest

	token := c.Query("token")
	videoId := c.Query("video_id")
	request.Token = token
	request.VideoId, _ = strconv.ParseInt(videoId, 10, 64)
	conn, err := grpc.Dial(":9003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to serve")
	}
	defer conn.Close()
	commentListClient := comment_grpc.NewCommentListServiceClient(conn)
	resp, err := commentListClient.ListComment(context.Background(), &request)
	c.JSON(http.StatusOK, &comment_grpc.DouYinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   resp.StatusMsg,
		CommentList: resp.CommentList,
	})
}
