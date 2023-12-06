package main

import (
	pb "PlanetMsg/idl/proto_gen/friend"
	"PlanetMsg/models"
	jwtInterceptor "PlanetMsg/pkg/interceptor"
	jwtUtil "PlanetMsg/pkg/jwt"
	"PlanetMsg/pkg/logger"
	"context"
	"net"

	"google.golang.org/grpc"
)

type FriendServer struct {
	pb.UnimplementedFriendServiceServer
}

func (*FriendServer) AddGroup(ctx context.Context, groupReq *pb.GroupReq) (res *pb.MsgRsp, err error) {
	// var userClaim *jwtUtil.UserClaim
	userClaim := ctx.Value("userClaim").(*jwtUtil.UserClaim)
	userID := userClaim.UserID
	msg, err := models.AddGroup(userID, groupReq.GroupName)
	if err != nil {
		return nil, err
	}
	res = &pb.MsgRsp{Msg: msg}
	return
}

// RouteGuide_ListFeaturesServer
func (*FriendServer) SearchUsers(nameReq *pb.NameReq, stream pb.FriendService_SearchUsersServer) error {
	users, err := models.SearchUsers(nameReq.Name)
	if err != nil {
		return err
	}
	for _, feature := range users {
		if err := stream.Send(&feature); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// TLS
	// c, err := credentials.NewServerTLSFromFile("../keys/server.pem", "../keys/server.key")
	// if err != nil {
	// 	logger.LogrusObj.Panic("获取TLS证书失败：%v", err)
	// }

	list, err := net.Listen("tcp", ":39002")
	if err != nil {
		logger.LogrusObj.Panic("端口请求失败：%v", err)
	}
	// server := grpc.NewServer(grpc.Creds(c))		// 开启TLS
	server := grpc.NewServer(
		grpc.UnaryInterceptor(jwtInterceptor.EnsureValidBasicCredentials),
	)
	pb.RegisterFriendServiceServer(server, new(FriendServer))
	err = server.Serve(list)
	if err != nil {
		logger.LogrusObj.Panic("用户服务上线失败：%v", err)
	}
}
