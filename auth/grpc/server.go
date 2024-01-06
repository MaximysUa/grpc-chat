package auth

import (
	"context"
	"google.golang.org/grpc"
	"grpc-chat/auth/proto"
)

type server struct {
	proto.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	proto.RegisterAuthServer(gRPC, &server{})
}

func (s *server) Login(ctx context.Context,
	req *proto.LoginRequest) (*proto.LoginResponse, error) {
	var resp proto.LoginResponse

	return &resp, nil
}
