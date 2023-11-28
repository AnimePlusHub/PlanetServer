package models

import (
	user_service "PlanetMsg/idl/proto_gen"
	"PlanetMsg/pkg/signalInfo"
	"PlanetMsg/pkg/util"
	// "gorm.io/gorm"
)

// 添加用户
func AddUser(user user_service.User) (string, error) {
	user.Pwd = util.MD5(user.Pwd)
	result := db.Create(&user)
	if result.Error != nil {
		return signalInfo.GetMsg(signalInfo.ERROR), result.Error
	}
	return signalInfo.GetMsg(signalInfo.SUCCESS), nil
}

// 查询用户
func GetUser(userId int32) *user_service.User {
	var user user_service.User
	db.Where("id=?", userId).Find(&user)
	return &user
}

// 通过邮箱查询用户
func GetUserByEmail(email string) *user_service.User {
	var user user_service.User
	db.Where("email=?", email).Find(&user)
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

// 根据Account查询用户id
func GetIdByAccount(account string) (userId int32) {
	row := db.Table("user").Where("account = ?", account).Select("id").Row()
	row.Scan(&userId)
	return
}

// 根据Email查询用户id
func GetIdByEmail(email string) (userId int32) {
	row := db.Table("user").Where("email = ?", email).Select("id").Row()
	row.Scan(&userId)
	return
}

// 根据Id查询用户密码
func GetPwdById(id int32) (pwd string) {
	row := db.Table("user").Where("id = ?", id).Select("pwd").Row()
	row.Scan(&pwd)
	return
}
