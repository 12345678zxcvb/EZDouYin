package routes

import (
	"awesomeProject4/Logic"
	"awesomeProject4/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	secretKey := []byte("zhang")
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", Logic.AuthMiddleware(secretKey), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", Logic.AuthMiddleware2(secretKey), controller.Publish)
	apiRouter.GET("/publish/list/", Logic.AuthMiddleware(secretKey), controller.PublishList)
	//
	//// extra apis - I
	apiRouter.POST("/favorite/action/", Logic.AuthMiddleware(secretKey), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", Logic.AuthMiddleware(secretKey), controller.FavoriteList)
	apiRouter.POST("/comment/action/", Logic.AuthMiddleware(secretKey), controller.CommentAction)
	apiRouter.GET("/comment/list/", Logic.AuthMiddleware(secretKey), controller.CommentList)
	//
	//// extra apis - II
	apiRouter.POST("/relation/action/", Logic.AuthMiddleware(secretKey), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", Logic.AuthMiddleware(secretKey), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", Logic.AuthMiddleware(secretKey), controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", Logic.AuthMiddleware(secretKey), controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageCheck)
	apiRouter.POST("/message/action/", controller.MessageAction)
}
