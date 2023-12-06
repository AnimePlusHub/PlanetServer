package models

type FriendGroup struct {
	ID        uint `gorm:"primaryKey"`
	UserId    int32
	GroupName string
}
