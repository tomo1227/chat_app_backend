package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomo1227/chat_app_backend/internal/adapter/ws/handler"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "hello")
	// })
	http.HandleFunc("/ws", handler.NewWebsocketHandler().Handle)
	port := "8080"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatalf("Serve Error: %v", err)
	}
}
