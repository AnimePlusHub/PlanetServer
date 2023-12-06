package main

import (
	pb "PlanetMsg/idl/proto_gen/sign_up"
	"PlanetMsg/mail"
	"PlanetMsg/models"
	jwtUtil "PlanetMsg/pkg/jwt"
	"PlanetMsg/pkg/logger"
	"PlanetMsg/pkg/signalInfo"
	"PlanetMsg/pkg/util"
	"PlanetMsg/redis"
	"context"
	"errors"
	"net"

	"google.golang.org/grpc"
)

type SignUpServer struct {
	pb.UnimplementedSignUpServiceServer
}

// 添加用户
func (u *SignUpServer) AddUser(ctx context.Context, user *pb.User) (*pb.MsgRsp, error) {
	defer func() {
		if err := recover(); err != nil {
			logger.LogrusObj.Printf("执行添加用户接口时出错:%v\n", err)
		}
	}()
	msg, err := models.AddUser(*user)
	return &pb.MsgRsp{Msg: msg}, err
}

func (u *SignUpServer) RequestValid(ctx context.Context, EmailReq *pb.EmailReq) (res *pb.MsgRsp, err error) {
	// 验证邮箱格式
	email := EmailReq.Email
	res = new(pb.MsgRsp)
	if util.ValidateEmail(email) == false {
		res.Msg = signalInfo.GetMsg(signalInfo.EMAIL_FOMAT_FALSE)
		err = errors.New("邮箱格式错误")
		return
	}
	// 生成随机六位数字
	validCode := util.GenerateRandom(6)
	// 插入Redis中，3分钟过期
	res.Msg, err = redis.Set(email, validCode, 3)
	// 向用户发送邮件
	mail.SendEmail(email, validCode)
	return
}

func (u *SignUpServer) CheckValidCode(ctx context.Context, checkEmailReq *pb.CheckEmailReq) (res *pb.MsgRsp, err error) {
	// 验证邮箱格式
	email := checkEmailReq.Email
	code := checkEmailReq.VaildCode
	res = new(pb.MsgRsp)
	if util.ValidateEmail(email) == false {
		res.Msg = signalInfo.GetMsg(signalInfo.EMAIL_FOMAT_FALSE)
		err = errors.New("邮箱格式错误")
		return
	}
	// 获取邮箱对应验证码
	rdsCode, resCode := redis.Get(email)
	// 判断是否获取value成功
	if resCode != signalInfo.GetMsg(signalInfo.REDIS_SUCCESS) {
		res.Msg = resCode
		return
	}
	// 判断验证码是否正确
	if code != rdsCode {
		res.Msg = signalInfo.GetMsg(signalInfo.VALIDATE_WRONG)
		return
	}
	res.Msg = signalInfo.GetMsg(signalInfo.VALIDATE_PASS)
	return
}

func (u *SignUpServer) Login(ctx context.Context, loginReq *pb.LoginReq) (res *pb.TokenRsp, err error) {
	res = new(pb.TokenRsp)
	loginName := loginReq.LoginName
	pwd := loginReq.Pwd
	var userId int32
	var account string
	// 判断用户的登录方式（邮箱 or Account）
	if util.ValidateEmail(loginName) { // 邮箱方式登录
		user := models.GetUserByEmail(loginName)
		userId = user.Id
		account = user.Account
	} else { // Account方式登录
		userId = models.GetIdByAccount(loginName)
		account = loginName
	}
	if userId == 0 { // 未找到用户
		res.Msg = signalInfo.GetMsg(signalInfo.ERROR_USER_NOT_FOUND)
		err = errors.New(signalInfo.GetMsg(signalInfo.ERROR_USER_NOT_FOUND))
		return
	}
	pwd = util.MD5(pwd)
	userPwd := models.GetPwdById(userId)
	if pwd != userPwd { // 密码错误
		res.Msg = signalInfo.GetMsg(signalInfo.ERROR_WRONG_PWD)
		err = errors.New(signalInfo.GetMsg(signalInfo.ERROR_WRONG_PWD))
		return
	}
	res.Atoken, res.Rtoken, err = jwtUtil.GenToken(userId, account)
	if err != nil {
		return nil, err
	}
	return
}

func main() {
	// TLS
	// c, err := credentials.NewServerTLSFromFile("../keys/server.pem", "../keys/server.key")
	// if err != nil {
	// 	logger.LogrusObj.Panic("获取TLS证书失败：%v", err)
	// }

	list, err := net.Listen("tcp", ":39000")
	if err != nil {
		logger.LogrusObj.Panic("端口请求失败：%v", err)
	}
	// server := grpc.NewServer(grpc.Creds(c))		// 开启TLS
	server := grpc.NewServer()
	pb.RegisterSignUpServiceServer(server, new(SignUpServer))
	err = server.Serve(list)
	if err != nil {
		logger.LogrusObj.Panic("用户服务上线失败：%v", err)
	}
}
