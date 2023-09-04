package model

type VideoTable struct {
	Id              int64 `gorm:"index"`
	UserInfoTableId int64
	User            UserInfoTable
	PlayUrl         string
	CoverUrl        string
	FavoriteCount   int64
	CommentCount    int64
	Title           string
	PublishTime     int64
}
