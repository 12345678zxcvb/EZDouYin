package relation_grpc

import (
	"awesomeProject4/Logic"
	"awesomeProject4/common"
	"awesomeProject4/common/model"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
	"net"
)

type followAction struct {
}

func (f followAction) GiveFollow(ctx context.Context, request *DouYinRelationActionRequest) (*DouYinRelationActionResponse, error) {
	token := request.Token
	secretKey := []byte("zhang")
	userName, err := Logic.ExtractUsernameFromToken(token, secretKey)
	if err != nil {
		fmt.Println(err)
	}
	var userInfoTable model.UserInfoTable
	db := common.GetDB()
	db.Take(&userInfoTable, "name = ?", userName)
	id := userInfoTable.ID
	relationTable := model.RelationTable{
		Follow:   request.ToUserId,
		Follower: userInfoTable.ID,
	}
	if request.ActionType == 1 {
		db.Create(&relationTable)
		userInfoTable.FollowCount = userInfoTable.FollowCount + 1
		db.Save(&userInfoTable)
		db.Take(&userInfoTable, request.ToUserId)
		userInfoTable.FollowerCount = userInfoTable.FollowerCount + 1
		db.Save(&userInfoTable)
	}
	if request.ActionType == 2 {
		db.Delete(&relationTable)
		db.Take(&userInfoTable, id)
		userInfoTable.FollowCount = userInfoTable.FollowCount - 1
		db.Save(&userInfoTable)
		db.Take(&userInfoTable, request.ToUserId)
		userInfoTable.FollowerCount = userInfoTable.FollowerCount - 1
		db.Save(&userInfoTable)
	}
	msg := "success"
	return &DouYinRelationActionResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
	}, nil
}

func (f followAction) mustEmbedUnimplementedFollowServiceServer() {

}

type followList struct {
}

func (f followList) FollowList(ctx context.Context, request *DouYinRelationFollowListRequest) (*DouYinRelationFollowListResponse, error) {
	var userInfoTable model.UserInfoTable
	db := common.GetDB()

	var relationTableList []model.RelationTable
	db.Find(&relationTableList, "follower = ?", request.UserId)
	var userList []*User
	for i := 0; i < len(relationTableList); i++ {
		followTable := relationTableList[i]
		db.Take(&userInfoTable, followTable.Follow)
		newUser := convertToUser(userInfoTable)
		userList = append(userList, newUser)
	}
	msg := "success"
	fmt.Println(userList)
	return &DouYinRelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		UserList:   userList,
	}, nil
}

func (f followList) mustEmbedUnimplementedFollowListServiceServer() {

}

type followerList struct {
}

func (f followerList) FollowerList(ctx context.Context, request *DouYinRelationFollowerListRequest) (*DouYinRelationFollowerListResponse, error) {

	var userInfoTable model.UserInfoTable
	db := common.GetDB()
	var relationTableList []model.RelationTable
	db.Find(&relationTableList, "follow = ?", request.UserId)
	var userList []*User
	for i := 0; i < len(relationTableList); i++ {
		followTable := relationTableList[i]
		db.Take(&userInfoTable, followTable.Follower)
		newUser := convertToUser(userInfoTable)
		userList = append(userList, newUser)
	}
	msg := "success"
	return &DouYinRelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		UserList:   userList,
	}, nil
}

func (f followerList) mustEmbedUnimplementedFollowerListServiceServer() {

}

type friendList struct {
}

func (f friendList) FriendList(ctx context.Context, request *DouYinRelationFriendListRequest) (*DouYinRelationFriendListResponse, error) {
	token := request.Token
	secretKey := []byte("zhang")
	userName, err := Logic.ExtractUsernameFromToken(token, secretKey)
	if err != nil {
		fmt.Println(err)
	}
	var userInfoTable model.UserInfoTable
	var friends []*FriendUser
	db := common.GetDB()
	db.Take(&userInfoTable, "name= ?", userName)
	friendSlice, err := FindFriendsList(db, userInfoTable.ID)
	if err != nil {
		msg := "success"
		return &DouYinRelationFriendListResponse{
			StatusCode: 0,
			StatusMsg:  &msg,
			UserList:   friends,
		}, nil
	}
	var userInfoTable1 model.UserInfoTable
	for i := 0; i < len(friendSlice); i++ {
		newInt := friendSlice[i]
		db.Take(&userInfoTable1, "id = ?", newInt)
		//newUser := convertToUser(userInfoTable1)
		text := "无"
		friend := FriendUser{
			Id:      userInfoTable1.ID,
			Name:    userInfoTable1.Name,
			MsgType: 0,
			Message: &text,
		}
		friends = append(friends, &friend)
	}
	msg := "success"
	return &DouYinRelationFriendListResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		UserList:   friends,
	}, nil
}

func (f friendList) mustEmbedUnimplementedFriendListServiceServer() {

}

func StartRelationServer() {
	grpcServer := grpc.NewServer()
	RegisterFollowServiceServer(grpcServer, &followAction{})
	RegisterFollowListServiceServer(grpcServer, &followList{})
	RegisterFollowerListServiceServer(grpcServer, &followerList{})
	RegisterFriendListServiceServer(grpcServer, &friendList{})
	listen, err := net.Listen("tcp", ":9004")
	if err != nil {
		log.Fatal("Failed to listen")
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("Failed to serve")
	}

}
func convertToUser(userinfotable model.UserInfoTable) *User {
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
	return &user
}
func FindFriendsList(db *gorm.DB, userID int64) ([]int64, error) {
	var friends []int64
	var relationList []model.RelationTable
	var relationInfo model.RelationTable
	db.Find(&relationList, "follow = ?", userID)
	for _, relation := range relationList {
		err := db.Take(&relationInfo, "follow = ?", relation.Follower).Error
		if err != nil {
			fmt.Println(err)
		}
		friends = append(friends, relationInfo.ID)
	}
	return friends, nil
}
