package feed_grpc

import (
	"awesomeProject4/common"
	"awesomeProject4/common/model"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type FeedStruct struct {
}

func (*FeedStruct) DouYinFeed(ctx context.Context, req *DouyinFeedRequest) (*DouyinFeedResponse, error) {
	currentTime := req.LatestTime
	db := common.GetDB()
	var videoTableList []model.VideoTable
	var videoTable model.VideoTable
	var videos []*Video
	if currentTime == 0 {
		db.Take(&videoTable, "id = ?", 1)
		newVideo := convertToVideo(&videoTable)
		videos = append(videos, newVideo)
		return &DouyinFeedResponse{
			StatusCode: 0,
			StatusMsg:  "success",
			VideoList:  videos,
			NextTime:   time.Now().Unix(),
		}, nil
	}
	err := db.Where("publish_time < ?", currentTime).Limit(20).Order("id DESC").Find(&videoTableList).Error
	if err != nil {
		db.Take(&videoTable, 1)
		finalTime := videoTable.PublishTime
		newVideo := convertToVideo(&videoTable)
		videos = append(videos, newVideo)

		return &DouyinFeedResponse{
			StatusCode: 0,
			StatusMsg:  "success",
			VideoList:  videos,
			NextTime:   finalTime,
		}, nil
	}
	for i := 0; i < len(videoTableList); i++ {
		video := videoTableList[i]
		newVideo := convertToVideo(&video)
		videos = append(videos, newVideo)
	}
	return &DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videos,
		NextTime:   videoTable.PublishTime,
	}, nil
}
func (*FeedStruct) mustEmbedUnimplementedFeedServiceServer() {

}

func StartFeedServer() {
	grpcServer := grpc.NewServer()
	feedStruct := &FeedStruct{}
	RegisterFeedServiceServer(grpcServer, feedStruct)
	listen, err := net.Listen("tcp", ":8787")
	if err != nil {
		log.Fatalf("Failed to listen")
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
	db.Take(&userinfotable, "id = ?", ID)
	user := User{
		Id:              userinfotable.ID,
		Name:            userinfotable.Name,
		FollowCount:     userinfotable.FollowCount,
		FollowerCount:   userinfotable.FollowerCount,
		IsFollow:        userinfotable.IsFollow,
		Avatar:          userinfotable.Avatar,
		BackgroundImage: userinfotable.BackgroundImage,
		Signature:       userinfotable.Signature,
		TotalFavorited:  userinfotable.TotalFavorite,
		WorkCount:       userinfotable.WorkCount,
		FavoriteCount:   userinfotable.FavoriteCount,
	}
	return user
}
