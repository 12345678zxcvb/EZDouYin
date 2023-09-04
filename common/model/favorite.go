package model

type FavoriteTable struct {
	Id              int64 `gorm:"index"`
	UserInfoTableId int64
	VideoTableId    int64
}
