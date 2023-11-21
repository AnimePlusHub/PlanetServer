package models

import (
	user_service "PlanetMsg/idl/proto_gen"
	"PlanetMsg/pkg/signalInfo"
	// "gorm.io/gorm"
)

// type User struct {
// 	ID            int    `gorm:"column:id"`
// 	Account       string `gorm:"column:account"`
// 	NickName      string `gorm:"column:nick_name"`
// 	Pwd           string
// 	Email         string
// 	Phone         string
// 	PicUrl        string
// 	BirthDay      time.Time
// 	CreateTime    time.Time
// 	Status        int
// 	LastLoginTime time.Time
// }

// 添加用户
func AddUser(user user_service.User) string {
	// user := User{
	// 	Account:       data["account"].(string),
	// 	NickName:      data["nick_name"].(string),
	// 	Pwd:           data["pwd"].(string),
	// 	Email:         data["email"].(string),
	// 	Phone:         data["phone"].(string),
	// 	PicUrl:        data["pic_url"].(string),
	// 	BirthDay:      util.ConvToDate(data["birth_day"].(string)),
	// 	CreateTime:    time.Now(),
	// 	Status:        data["status"].(int),
	// 	LastLoginTime: util.ConvToDateTime(data["last_login_time"].(string)),
	// }
	db.Create(&user)
	if user.Id > 0 {
		return signalInfo.GetMsg(signalInfo.SUCCESS)
	}
	return signalInfo.GetMsg(500)
}

// 查询用户
func GetUser(userId int32) *user_service.User {
	var user user_service.User
	db.Where("id=?", userId).Find(&user)
	return &user
}

// 更新用户单个字段
func UpdateUserSingle(updateKey string, user user_service.User) string {
	result := db.Model(user).Select(updateKey).Updates(user)
	if result.Error != nil {
		return result.Error.Error()
	}
	return signalInfo.GetMsg(signalInfo.SUCCESS)
}

// 更新用户（非0）字段信息
func UpdateUser(user user_service.User) string {
	result := db.Model(&user).Updates(user)
	if result.Error != nil {
		return result.Error.Error()
	}
	return signalInfo.GetMsg(signalInfo.SUCCESS)
}
