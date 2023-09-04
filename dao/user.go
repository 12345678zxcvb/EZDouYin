package dao

import (
	"awesomeProject4/common"
	"awesomeProject4/common/model"
)

func CreateTable(username, password string) {
	userInfoTable := model.UserInfoTable{
		Name:     username,
		PassWord: password,
	}
	db := common.GetDB()
	db.Create(&userInfoTable)
}
