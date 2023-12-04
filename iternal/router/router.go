package router

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"grpc-chat/api/proto"
	"grpc-chat/iternal/handler"
)

func New(ctx context.Context, client proto.ChatClient) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Post("/api/send", handler.SendMessage(ctx, client))
	router.Post("/api/receive", handler.ReceiveMessage(ctx, client))
	return router
}
