package controller

import (
	"awesomeProject4/douyin_grpc/message_grpc"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
)

func MessageAction(c *gin.Context) {
	var req message_grpc.DouYinMessageActionRequest
	req.Content = c.Query("content")
	actionType := c.Query("action_type")
	newInt, _ := strconv.ParseInt(actionType, 10, 64)
	req.ActionType = int32(newInt)
	toUserId := c.Query("to_user_id")
	req.ToUserId, _ = strconv.ParseInt(toUserId, 10, 64)
	req.Token = c.Query("token")
	conn, err := grpc.Dial(":9005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to Dial",
		})
	}
	defer conn.Close()
	chatClient := message_grpc.NewMessageActionClient(conn)
	resp, err := chatClient.Chat(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to Chat",
		})
	}
	c.JSON(http.StatusOK, message_grpc.DouYinMessageActionResponse{
		StatusCode: 0,
		StatusMsg:  resp.StatusMsg,
	})
}
func MessageCheck(c *gin.Context) {
	var req message_grpc.DouYinMessageChatRequest
	toUserId := c.Query("to_user_id")
	req.ToUserId, _ = strconv.ParseInt(toUserId, 10, 64)
	req.Token = c.Query("token")
	newTime := c.Query("pre_msg_time")
	req.PreMsgTime, _ = strconv.ParseInt(newTime, 10, 64)
	conn, err := grpc.Dial(":9005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to Dial",
		})
	}
	defer conn.Close()
	messageClient := message_grpc.NewChatServiceClient(conn)
	resp, err := messageClient.CheckMessage(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to Chat",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":  0,
		"status_msg":   resp.StatusMsg,
		"message_list": resp.MessageList,
	})
}
