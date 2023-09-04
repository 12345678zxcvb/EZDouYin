package message_grpc

import (
	"awesomeProject4/Logic"
	"awesomeProject4/common"
	"awesomeProject4/common/model"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type chat struct {
}

func (c chat) Chat(ctx context.Context, request *DouYinMessageActionRequest) (*DouYinMessageActionResponse, error) {
	secretKey := []byte("zhang")
	userName, _ := Logic.ExtractUsernameFromToken(request.Token, secretKey)
	var userInfoTable model.UserInfoTable
	db := common.GetDB()
	db.Take(&userInfoTable, "name = ?", userName)
	messageInfo := model.MessageTable{
		FromUserID: userInfoTable.ID,
		ToUserID:   request.ToUserId,
		Content:    request.Content,
		CreateTime: time.Now().Unix(),
	}
	fmt.Println(messageInfo)
	db.Create(&messageInfo)
	msg := "success"
	return &DouYinMessageActionResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
	}, nil
}

func (c chat) mustEmbedUnimplementedMessageActionServer() {

}

type checkMessage struct {
}

func (c checkMessage) CheckMessage(ctx context.Context, request *DouYinMessageChatRequest) (*DouYinMessageChatResponse, error) {
	preTime := request.PreMsgTime
	secretKey := []byte("zhang")
	userName, _ := Logic.ExtractUsernameFromToken(request.Token, secretKey)
	var userInfoTable model.UserInfoTable
	db := common.GetDB()
	db.Take(&userInfoTable, "name = ?", userName)
	var messages []*Message
	var messageList []model.MessageTable
	err := db.Where("(from_user_id = ? AND to_user_id = ? OR to_user_id = ? AND from_user_id = ?) AND create_time > ?",
		userInfoTable.ID, request.ToUserId, request.ToUserId, userInfoTable.ID, preTime).
		Find(&messageList).Error
	if err == nil {
		for _, messageInfo := range messageList {
			newMessage := convertToMessage(messageInfo)
			messages = append(messages, newMessage)
		}
		msg := "success"
		return &DouYinMessageChatResponse{
			StatusCode:  0,
			StatusMsg:   &msg,
			MessageList: messages,
		}, nil
	}
	msg := "success"
	return &DouYinMessageChatResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
	}, nil
}

func (c checkMessage) mustEmbedUnimplementedChatServiceServer() {

}

func StartMessageServer() {
	grpcServer := grpc.NewServer()
	RegisterMessageActionServer(grpcServer, &chat{})
	RegisterChatServiceServer(grpcServer, &checkMessage{})
	listen, err := net.Listen("tcp", ":9005")
	if err != nil {
		log.Fatalf("Failed to listen")
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve")
	}
}
func convertToMessage(messageTable model.MessageTable) *Message {
	newTime := strconv.FormatInt(messageTable.CreateTime, 10)
	message := &Message{
		Id:         messageTable.ID,
		FromUserId: messageTable.FromUserID,
		ToUserId:   messageTable.ToUserID,
		Content:    messageTable.Content,
		CreateTime: &newTime,
	}
	return message
}
