package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// オリジンの確認を行わない場合は、常に true を返す
		return true
	},
}

func handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Client Connected")

	for {
		// メッセージの読み込み
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// 受信したメッセージを出力
		fmt.Printf("Received Message: %s\n", p)

		// 受信したメッセージをそのまま返す
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocketConnection)

	fmt.Println("WebSocket server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
