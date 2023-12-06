package main

import (
	// tpb "PlanetMsg/idl/proto_gen/sign_up"
	pb "PlanetMsg/idl/proto_gen/user"
	"PlanetMsg/models"
	jwtInterceptor "PlanetMsg/pkg/interceptor"
	"PlanetMsg/pkg/logger"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

// 通过id查询用户信息
func (u *UserServer) GetUserInfo(ctx context.Context, request *pb.IdReq) (*pb.User, error) {
	defer func() {
		if err := recover(); err != nil {
			logger.LogrusObj.Printf("执行查询用户接口时出错:%v\n", err)
		}
	}()
	userId := request.UserId
	user, err := models.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
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

// 更新用户信息
func (u *UserServer) UpdateUser(ctx context.Context, user *pb.User) (*pb.MsgRsp, error) {
	defer func() {
		if err := recover(); err != nil {
			logger.LogrusObj.Printf("执行更新用户接口时出错:%v\n", err)
		}
	}()
	msg := models.UpdateUser(*user)
	return &pb.MsgRsp{Msg: msg}, nil
}

func main() {
	// TLS
	// c, err := credentials.NewServerTLSFromFile("../keys/server.pem", "../keys/server.key")
	// if err != nil {
	// 	logger.LogrusObj.Panic("获取TLS证书失败：%v", err)
	// }

	list, err := net.Listen("tcp", ":39001")
	if err != nil {
		logger.LogrusObj.Panic("端口请求失败：%v", err)
	}
	// server := grpc.NewServer(grpc.Creds(c))		// 开启TLS
	server := grpc.NewServer(
		grpc.UnaryInterceptor(jwtInterceptor.EnsureValidBasicCredentials),
	)
	pb.RegisterUserServiceServer(server, new(UserServer))
	err = server.Serve(list)
	if err != nil {
		logger.LogrusObj.Panic("用户服务上线失败：%v", err)
	}
}
