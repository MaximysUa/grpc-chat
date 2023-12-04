package main

import (
	"google.golang.org/grpc"
	server "grpc-chat/api/grpc-server"
	api "grpc-chat/api/proto"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := grpc.NewServer()
	srv := server.GRPCServer{}
	api.RegisterChatServer(s, srv)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err = s.Serve(listen)
		if err != nil {
			log.Fatal(err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	log.Println("Stopping server")
	s.GracefulStop()
}
