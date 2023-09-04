package comment_grpc

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

type giveComment struct {
}

func (c *giveComment) mustEmbedUnimplementedCommentServiceServer() {

}

func (c *giveComment) GiveComment(ctx context.Context, request *DouYinCommentActionRequest) (*DouYinCommentActionResponse, error) {
	token := request.Token
	secretKey := []byte("zhang")
	userName, err := Logic.ExtractUsernameFromToken(token, secretKey)
	if err != nil {
		log.Fatal("Failed to get username")
	}
	var userInfoTable model.UserInfoTable
	db := common.GetDB()
	db.Take(&userInfoTable, "name = ?", userName)
	commentTable := model.CommentTable{
		UserInfoTableID: userInfoTable.ID,
		VideoTableID:    request.VideoId,
		CommentText:     *request.CommentText,
		Time:            Logic.GetCurrentTime(),
	}
	if request.ActionType == 1 {
		db.Create(&commentTable)
		var videoInfoTable model.VideoTable
		db.Take(&videoInfoTable, request.VideoId)
		videoInfoTable.CommentCount = videoInfoTable.CommentCount + 1
		db.Save(&videoInfoTable)
	}
	if request.ActionType == 2 {
		db.Delete(&commentTable)
		var videoInfoTable model.VideoTable
		db.Take(&videoInfoTable, request.VideoId)
		videoInfoTable.CommentCount = videoInfoTable.CommentCount - 1
		db.Save(&videoInfoTable)
	}
	msg := "success"
	newComment := convertToComment(commentTable)
	return &DouYinCommentActionResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		Comment:    newComment,
	}, nil
}

type commentList struct {
}

func (c commentList) mustEmbedUnimplementedCommentListServiceServer() {

}
func (c commentList) ListComment(ctx context.Context, request *DouYinCommentListRequest) (*DouYinCommentListResponse, error) {
	token := request.Token
	secretKey := []byte("zhang")
	userName, err := Logic.ExtractUsernameFromToken(token, secretKey)
	if err != nil {
		fmt.Println(err)
	}
	var userInfoTable model.UserInfoTable
	db := common.GetDB()
	db.Take(&userInfoTable, "name = ?", userName)
	var comments []*Comment
	var commentTableList []model.CommentTable
	db.Find(&commentTableList, "video_table_id = ?", request.VideoId)
	for _, commentTable := range commentTableList {
		newComment := convertToComment(commentTable)
		comments = append(comments, newComment)
	}
	if err != nil {
		log.Fatal(err)
	}
	msg := "success"
	return &DouYinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   &msg,
		CommentList: comments,
	}, nil
}
func StartCommentServer() {
	grpcServer := grpc.NewServer()
	RegisterCommentServiceServer(grpcServer, &giveComment{})
	RegisterCommentListServiceServer(grpcServer, &commentList{})
	listen, err := net.Listen("tcp", ":9003")
	if err != nil {
		log.Fatal("Failed to listen")
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve")
	}
}
func convertToComment(commentTable model.CommentTable) *Comment {
	var userInfoTable model.UserInfoTable
	db := common.GetDB()
	db.Take(&userInfoTable, commentTable.UserInfoTableID)
	newUser := convertToUser(userInfoTable)
	comment := Comment{
		Id:         commentTable.ID,
		Content:    commentTable.CommentText,
		CreateDate: commentTable.Time,
		User:       &newUser,
	}
	return &comment
}
func convertToUser(userinfotable model.UserInfoTable) User {
	user := User{
		Id:              userinfotable.ID,
		Name:            userinfotable.Name,
		FollowCount:     &userinfotable.FollowCount,
		FollowerCount:   &userinfotable.FollowerCount,
		IsFollow:        userinfotable.IsFollow,
		Avatar:          &userinfotable.Avatar,
		BackgroundImage: &userinfotable.BackgroundImage,
		Signature:       &userinfotable.Signature,
		TotalFavorited:  &userinfotable.TotalFavorite,
		WorkCount:       &userinfotable.WorkCount,
		FavoriteCount:   &userinfotable.FavoriteCount,
	}
	return user
}
