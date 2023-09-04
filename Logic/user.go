package Logic

import (
	"awesomeProject4/common"
	"awesomeProject4/common/model"
	"awesomeProject4/dao"
)

func LogInVerification(username string, password string) string {
	var str string
	var userInfoTable model.UserInfoTable
	db := common.GetDB()
	db.Take(&userInfoTable, "Name = ?", username)
	if userInfoTable.PassWord != password {
		str = "Password is wrong"
	}
	return str
}
func RegisterVerification(username string, password string) string {
	var str string
	var userInfoTable model.UserInfoTable
	db := common.GetDB()
	if len(username) < 6 {
		str = "The length of username cannot be less than 6 "
	}
	if len(password) > 32 {
		str = "The length of password cannot be large than 32"
	}
	err := db.Take(&userInfoTable, "name = ?", username).Error
	if err == nil {
		str = "User is already exist"
	}
	if str == "" {
		dao.CreateTable(username, password)
	}
	return str
}
