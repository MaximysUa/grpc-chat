package handler

import (
	"context"
	"github.com/go-chi/render"
	"grpc-chat/api/proto"
	"log"
	"net/http"
)

func SendMessage(ctx context.Context, client proto.ChatClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req proto.Message
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, "Error")
			return
		}
		message, err := client.SendMessage(ctx, &req)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, "Error")
			return
		}
		log.Println(message)
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, message)
	}
}

func ReceiveMessage(ctx context.Context, client proto.ChatClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req proto.ReceiveRequest
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, "Error")
			return
		}
		message, err := client.ReceiveMessage(ctx, &req)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, "Error")
			return
		}
		log.Println(message)
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, message)
	}
}
