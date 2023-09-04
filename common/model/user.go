package model

type UserInfoTable struct {
	ID              int64 `gorm:"index"`
	Name            string
	PassWord        string
	FollowCount     int64
	FollowerCount   int64
	IsFollow        bool
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorite   int64
	WorkCount       int64
	FavoriteCount   int64
	Video           []VideoTable
}
