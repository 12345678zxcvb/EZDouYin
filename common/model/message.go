package model

type MessageTable struct {
	ID         int64 `gorm:"index"`
	FromUserID int64
	ToUserID   int64
	Content    string
	CreateTime int64
}
