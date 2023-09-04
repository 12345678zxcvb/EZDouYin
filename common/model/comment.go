package model

type CommentTable struct {
	ID              int64 `gorm:"index"`
	UserInfoTableID int64
	VideoTableID    int64
	CommentText     string
	Time            string
}
