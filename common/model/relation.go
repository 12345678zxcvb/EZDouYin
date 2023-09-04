package model

type RelationTable struct {
	ID       int64 `gorm:"index"`
	Follow   int64
	Follower int64
}
