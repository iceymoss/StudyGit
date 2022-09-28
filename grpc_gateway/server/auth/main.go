package main

import (
	authpb "grpc_gateway/server/auth/api/gen/v1"
	"log"
	"net"

	"grpc_gateway/server/auth/auth"

	"google.golang.org/grpc"
)

//service端
func main() {
	conn, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	//实例化server
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{})
	s.Serve(conn)
}
