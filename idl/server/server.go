package main

import (
	pb "PlanetMsg/idl/proto_gen"
	"PlanetMsg/mail"
	"PlanetMsg/models"
	"PlanetMsg/pkg/signalInfo"
	"PlanetMsg/pkg/util"
	"PlanetMsg/redis"
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

// 添加用户
func (u *UserServer) AddUser(ctx context.Context, user *pb.User) (*pb.MsgRsp, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("执行添加用户接口时出错:%v\n", err)
		}
	}()
	msg := models.AddUser(*user)
	return &pb.MsgRsp{Msg: msg}, nil
}

// 更新用户信息
func (u *UserServer) UpdateUser(ctx context.Context, user *pb.User) (*pb.MsgRsp, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("执行更新用户接口时出错:%v\n", err)
		}
	}()
	msg := models.UpdateUser(*user)
	return &pb.MsgRsp{Msg: msg}, nil
}

// 根据字段更新用户信息
func (u *UserServer) UpdateUserSingle(ctx context.Context, request *pb.UpdateSingleReq) (*pb.MsgRsp, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("执行更新用户字段接口时出错:%v\n", err)
		}
	}()
	msg := models.UpdateUserSingle(request.UpdateKey, *request.User)
	return &pb.MsgRsp{Msg: msg}, nil
}

// 通过id查询用户信息
func (u *UserServer) GetUserInfo(ctx context.Context, request *pb.IdReq) (*pb.User, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("执行查询用户接口时出错:%v\n", err)
		}
	}()
	userId := request.UserId
	user := models.GetUser(userId)
	return user, nil
}

func (u *UserServer) RequestValid(ctx context.Context, EmailReq *pb.EmailReq) (res *pb.MsgRsp, err error) {
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

func (u *UserServer) CheckValidCode(ctx context.Context, checkEmailReq *pb.CheckEmailReq) (res *pb.MsgRsp, err error) {
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

func main() {
	list, err := net.Listen("tcp", ":39001")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, new(UserServer))
	err = server.Serve(list)
	if err != nil {
		panic(err)
	}
}
