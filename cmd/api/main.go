package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomo1227/chat_app_backend/internal/adapter/ws/handler"
	"github.com/tomo1227/chat_app_backend/internal/domain"
)

func main() {
	hub := domain.NewHub()
	go hub.RunLoop()

	// http.HandleFunc("/ws", handler.NewWebsocketHandler(hub).Handle)
	http.HandleFunc("/ws", handler.NewWebsocketHandler().Handle)
	port := "8080"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatalf("Serve Error: %v", err)
	}
}
