package auth

import (
	"fmt"
	authpb "grpc_gateway/server/auth/api/gen/v1"
	"strconv"

	"golang.org/x/net/context"
)

type Service struct {
}

//后台事件处理方法
func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	exp, err := strconv.Atoi(req.Code)
	if err != nil {
		fmt.Println(err)
	}

	return &authpb.LoginResponse{
		AccssToken: "恭喜您成功获取到令牌",
		ExpiresIn:  int32(exp),
	}, nil
}
