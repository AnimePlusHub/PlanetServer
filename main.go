package main

import (
	"PlanetMsg/mail"
)

func main() {
	// 新增一个用户
	// data := make(map[string]interface{})
	// data["account"] = "account123"
	// data["nick_name"] = "TechAoba"
	// data["pwd"] = "pwd123456"
	// data["email"] = "techaboa@163.com"
	// data["phone"] = "12451344251"
	// data["pic_url"] = "xxx.png"
	// data["birth_day"] = "2000-02-18"
	// data["status"] = 0
	// data["last_login_time"] = "2021-05-11 14:31:19"

	// models.AddUser(data)
	// code := signalInfo.SUCCESS
	// fmt.Println(code)

	// 查询新用户
	// user := models.GetUser(1)
	// if user != nil {
	// 	fmt.Printf("%+v\b", *user)
	// } else {
	// 	fmt.Println("无结果")
	// }

	// 更新用户信息
	// user := models.User{
	// 	ID:      1,
	// 	Account: "newAcc1",
	// 	Email:   "newEml",
	// }

	// models.UpdateUserSingle("Account", user)

	// sourceSlice := []int{1, 2, 3, 4, 5}
	// dstSlice := make([]int, len(sourceSlice))
	// numberCopied := copy(dstSlice, sourceSlice)
	// fmt.Println("Copied Numbers: ", numberCopied)
	// fmt.Println("dstSlice: ", dstSlice)
	// sourceSlice[1] = 100
	// fmt.Println("sourceSlice: ", sourceSlice)
	// fmt.Println("dstSlice: ", dstSlice)

	// redis.Rds()

	mail.SendEmail("1079334524@qq.com")
}
