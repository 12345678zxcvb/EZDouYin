package main

import (
	"awesomeProject4/common"
	"awesomeProject4/douyin_grpc/comment_grpc"
	"awesomeProject4/douyin_grpc/favorite_grpc"
	"awesomeProject4/douyin_grpc/feed_grpc"
	"awesomeProject4/douyin_grpc/message_grpc"
	"awesomeProject4/douyin_grpc/publish_grpc"
	"awesomeProject4/douyin_grpc/relation_grpc"
	"awesomeProject4/douyin_grpc/user_grpc"
	"awesomeProject4/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	router := gin.Default()
	routes.InitRouter(router)
	go user_grpc.StartGRPCServer()
	go feed_grpc.StartFeedServer()
	go publish_grpc.StartPublishServer()
	go favorite_grpc.StartFavoriteServer()
	go comment_grpc.StartCommentServer()
	go relation_grpc.StartRelationServer()
	go message_grpc.StartMessageServer()
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err, "  Failed to start. Please try again!")
	}
}
