package main

import (
	proto "grpc_gateway/server/trip/api"
	"grpc_gateway/server/trip/trip"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	conn, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal("监听端口失败", err)
	}

	s := grpc.NewServer()
	proto.RegisterTripServiceServer(s, &trip.TripService{})
	s.Serve(conn)
}
