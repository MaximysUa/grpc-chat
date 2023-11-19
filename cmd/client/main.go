package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	api "grpc-chat/api/proto"
	"log"
)

func main() {
	dial, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	ctx := context.TODO()
	client := api.NewChatClient(dial)
	m, err := client.SendMessage(ctx, &api.Message{
		Sender: &api.User{
			Id:        1,
			FirstName: "Ihor",
			LastName:  "Adamov",
			Age:       20,
		},
		Receiver: &api.User{
			Id:        2,
			FirstName: "adsda",
			LastName:  "dddww",
			Age:       33,
		},
		Text: "Hi",
	})
	if err != nil {
		return
	}
	log.Println(m)
	message, err := client.ReceiveMessage(ctx, &api.ReceiveRequest{Id: 2})
	if err != nil {
		return
	}
	log.Println(message.Mess)
}
