package main

import (
	"context"
	"fmt"
	authpb "grpc_gateway/server/auth/api/gen/v1"
	proto "grpc_gateway/server/trip/api"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func GateWay() {
	c := context.Background()          //生成上下文
	c, cancel := context.WithCancel(c) //生成有cancel功能
	defer cancel()

	//new a mux for server
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			EnumsAsInts: true, //status
			OrigName:    true, //命名
		},
	))

	//连接auth服务
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		c,
		mux,
		":8089", //连接grpc服务
		[]grpc.DialOption{grpc.WithInsecure()},
	)

	//连接trip服务
	err = proto.RegisterTripServiceHandlerFromEndpoint(
		c,
		mux,
		":8001",
		[]grpc.DialOption{grpc.WithInsecure()},
	)

	if err != nil {
		log.Fatalf("断开连接: %v", err)
	}

	//开启http服务，对外暴露端口
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}

}
func main() {
	fmt.Println("开始")
	GateWay()
}
