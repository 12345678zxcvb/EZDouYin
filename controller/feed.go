package controller

import (
	"awesomeProject4/douyin_grpc/feed_grpc"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
	"time"
)

func Feed(c *gin.Context) {
	var lastTime = c.Query("latest_time")

	//timeNum, err := strconv.ParseInt(lastTime, 10, 64)
	var req feed_grpc.DouyinFeedRequest
	req.LatestTime, _ = strconv.ParseInt(lastTime, 10, 64)
	conn, err := grpc.Dial(":8787", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to connection",
		})
	}
	defer conn.Close()
	feedClient := feed_grpc.NewFeedServiceClient(conn)
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	resp, err := feedClient.DouYinFeed(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Feed"})
		return
	}
	c.JSON(http.StatusOK, feed_grpc.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  resp.VideoList,
		NextTime:   time.Now().Unix(),
	})
}
