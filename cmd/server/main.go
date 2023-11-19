package main

import (
	"google.golang.org/grpc"
	server "grpc-chat/api/grpc-server"
	api "grpc-chat/api/proto"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := server.GRPCServer{}
	api.RegisterChatServer(s, srv)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
