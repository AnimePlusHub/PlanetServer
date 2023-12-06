package models

import (
	friend_service "PlanetMsg/idl/proto_gen/friend"
	"PlanetMsg/pkg/signalInfo"
	// "gorm.io/gorm"
)

// 添加好友分组
func AddGroup(userId int32, groupName string) (string, error) {
	group := FriendGroup{0, userId, groupName}
	result := db.Create(&group)
	if result.Error != nil {
		return "", result.Error
	}
	return signalInfo.GetMsg(signalInfo.SUCCESS), nil
}

// 根据account或昵称查找多个用户
func SearchUsers(name string) (users []friend_service.User, err error) {
	result := db.Where("account LIKE ? OR nick_name LIKE ?", name+"%", name+"%").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
