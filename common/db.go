package common

import (
	"awesomeProject4/common/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	username := "root"            //账号
	password := "clx2575167@0537" //密码
	host := "127.0.0.1"           //数据库地址，可以是Ip或者域名
	port := 3306                  //数据库端口
	Dbname := "ezdouyin"          //数据库名
	timeout := "10s"              //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open("mysql", dsn)
	db.AutoMigrate(&model.UserInfoTable{})
	db.AutoMigrate(&model.VideoTable{})
	db.AutoMigrate(&model.FavoriteTable{})
	db.AutoMigrate(&model.CommentTable{})
	db.AutoMigrate(&model.RelationTable{})
	db.AutoMigrate(&model.MessageTable{})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}
