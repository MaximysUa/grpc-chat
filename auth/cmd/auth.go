package auth

import (
	"google.golang.org/grpc"
	auth "grpc-chat/auth/grpc"
)

type Auth struct {
	gRPCServer *grpc.Server
	port       int
}

func New(port int) (*Auth, error) {
	//proto.RegisterAuthServer(grpc.NewServer(), auth.Server{})
	auth.Register(grpc.NewServer())
}
