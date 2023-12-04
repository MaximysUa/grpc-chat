package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	api "grpc-chat/api/proto"
	"grpc-chat/iternal/router"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	dial, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	ctx := context.TODO()
	client := api.NewChatClient(dial)

	mux := router.New(ctx, client)
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.Serve(listen))

}
