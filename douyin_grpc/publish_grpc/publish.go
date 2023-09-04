package publish_grpc

import (
	"awesomeProject4/common"
	"awesomeProject4/common/model"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type publishAction struct {
}

func (a *publishAction) mustEmbedUnimplementedActionPublishServer() {

	panic("implement me")
}

func (*publishAction) PublishAction(ctx context.Context, req *DouYinPublishActionRequest) (*DouYinPublishActionResponse, error) {
	msg := "success"
	return &DouYinPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
	}, nil
}

type publishList struct {
}

func (*publishList) mustEmbedUnimplementedListPublishServer() {
	//TODO implement me
	panic("implement me")
}

func (*publishList) PublishList(ctx context.Context, req *DouYinPublishListRequest) (*DouYinPublishListResponse, error) {
	var videos []*Video
	db := common.GetDB()
	var videoTables []model.VideoTable

	db.Where("user_info_table_id IN (?)", []int64{req.UserId}).Find(&videoTables)
	for _, videoTable := range videoTables {
		video := convertToVideo(&videoTable)
		videos = append(videos, video)
	}
	msg := "success"
	return &DouYinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		VideoList:  videos,
	}, nil
}
func StartPublishServer() {
	grpcServer := grpc.NewServer()
	RegisterActionPublishServer(grpcServer, &publishAction{})
	RegisterListPublishServer(grpcServer, &publishList{})
	listen, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
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
