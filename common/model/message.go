package model

type MessageTable struct {
	ID         int64 `gorm:"index"`
	ToUserID   int64
	FromUserID int64
	Content    string
	CreateTime int64
}
