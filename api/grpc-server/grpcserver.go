package server

import (
	"context"
	"grpc-chat/api/database"
	api "grpc-chat/api/proto"
)

var db = make([]api.Message, 0, 5)

type GRPCServer struct {
	api.UnimplementedChatServer
}

func (s GRPCServer) SendMessage(ctx context.Context, mess *api.Message) (*api.SendResponse, error) {
	var resp api.SendResponse
	db = append(db, *mess)

	resp.Status = "ok"
	return &resp, nil
}

func (s GRPCServer) ReceiveMessage(ctx context.Context, mess *api.ReceiveRequest) (*api.ReceiveResponse, error) {
	var resp api.ReceiveResponse
	for _, i := range db {

		if i.Receiver.Id == mess.Id {
			resp.Mess = append(resp.Mess, &i)
		}
	}
	return &resp, nil
}

func (s GRPCServer) New(db *database.DB) (*GRPCServer, error) {
	return &s, nil
}
