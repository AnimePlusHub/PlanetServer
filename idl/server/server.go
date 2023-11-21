package main

import (
	pb "PlanetMsg/idl/proto_gen"
	"PlanetMsg/models"
	"context"
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
