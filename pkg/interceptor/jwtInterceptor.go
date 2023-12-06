package jwtInterceptor

import (
	jwtUtil "PlanetMsg/pkg/jwt"
	"PlanetMsg/pkg/signalInfo"
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
)

func EnsureValidBasicCredentials(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	authStr := md["authorization"]
	if len(authStr) == 0 { // 缺少token
		return nil, errors.New(signalInfo.GetMsg(signalInfo.MISS_TOKEN))
	}
	tokenString := strings.TrimPrefix(authStr[0], "Bearer ")

	userClaim, err := jwtUtil.VerifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	// logger.LogrusObj.Println(claims.ID)
	ctx = context.WithValue(ctx, "userClaim", userClaim)
	md, ok = metadata.FromIncomingContext(ctx)

	return handler(ctx, req)
}
