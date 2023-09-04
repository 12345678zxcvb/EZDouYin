package favorite_grpc

import (
	"awesomeProject4/Logic"
	"awesomeProject4/common"
	"awesomeProject4/common/model"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type favoriteAction struct {
}

func (f favoriteAction) GiveFavoriteService(ctx context.Context, request *DouYinFavoriteActionRequest) (*DouYinFavoriteActionResponse, error) {
	secretKey := []byte("zhang")
	userName, _ := Logic.ExtractUsernameFromToken(request.Token, secretKey)
	db := common.GetDB()
	var userInfoTable model.UserInfoTable
	db.Take(&userInfoTable, "Name = ?", userName)
	id := userInfoTable.ID
	if request.ActionType == 1 {
		favoriteTable := model.FavoriteTable{
			UserInfoTableId: id,
			VideoTableId:    request.VideoId,
		}
		db.Create(&favoriteTable)
		var videoInfoTable model.VideoTable
		db.Take(&videoInfoTable, request.VideoId)
		videoInfoTable.FavoriteCount = videoInfoTable.FavoriteCount + 1
		db.Save(&videoInfoTable)
		db.Take(&userInfoTable, videoInfoTable.UserInfoTableId)
		userInfoTable.FavoriteCount = userInfoTable.FavoriteCount + 1
		db.Save(&userInfoTable)
	}
	if request.ActionType == 2 {
		var favoriteTable model.FavoriteTable
		favoriteTable.UserInfoTableId = id
		favoriteTable.VideoTableId = request.VideoId
		db.Take(&favoriteTable)
		db.Delete(&favoriteTable)
		var videoInfoTable model.VideoTable
		db.Take(&videoInfoTable, request.VideoId)
		videoInfoTable.FavoriteCount = videoInfoTable.FavoriteCount - 1
		db.Save(&videoInfoTable)
		db.Take(&userInfoTable, videoInfoTable.UserInfoTableId)
		userInfoTable.FavoriteCount = userInfoTable.FavoriteCount - 1
		db.Save(&userInfoTable)
	}

	msg := "success"
	return &DouYinFavoriteActionResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
	}, nil
}

func (f favoriteAction) mustEmbedUnimplementedFavoriteActionServer() {

	panic("implement me")
}

type favoriteList struct {
}

func (f favoriteList) FavoriteListService(ctx context.Context, request *DouYinFavoriteListRequest) (*DouYinFavoriteListResponse, error) {

	var userInfoTable model.UserInfoTable
	db := common.GetDB()
	db.Take(&userInfoTable, request.UserId)
	var favoriteTableList []model.FavoriteTable
	db.Find(&favoriteTableList, "user_info_table_id = ?", userInfoTable.ID)

	var videoList []*Video
	for _, favoriteTable := range favoriteTableList {
		var videoTable model.VideoTable
		db.Take(&videoTable, favoriteTable.VideoTableId)
		newVideo := convertToVideo(&videoTable)
		videoList = append(videoList, newVideo)
	}

	msg := "success"
	return &DouYinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		VideoList:  videoList,
	}, nil

}

func (f favoriteList) mustEmbedUnimplementedFavoriteListServer() {

	panic("implement me")
}

func StartFavoriteServer() {
	grpcServer := grpc.NewServer()

	RegisterFavoriteActionServer(grpcServer, &favoriteAction{})
	RegisterFavoriteListServer(grpcServer, &favoriteList{})
	listen, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatal("Failed to listen")
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve")
	}

}
func convertToVideo(videoTable *model.VideoTable) *Video {

	author := convertToUser(videoTable.UserInfoTableId)

	convertVideo := &Video{
		Id:            videoTable.Id,
		Author:        &author,
		PlayUrl:       videoTable.PlayUrl,
		CoverUrl:      videoTable.CoverUrl,
		FavoriteCount: videoTable.FavoriteCount,
		CommentCount:  videoTable.CommentCount,
		Title:         videoTable.Title,
	}
	return convertVideo

}
func convertToUser(ID int64) User {
	db := common.GetDB()
	var userinfotable model.UserInfoTable
	db.Take(&userinfotable, ID)
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
