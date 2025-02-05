package main

import (
	"log"
	"net/http"

	"github.com/fullavatar/realTimeChat/websocket_chat"
)

func main() {
	server := websocket_chat.NewWebSocketServer()

	http.Handle("/ws", server)

	port := "5000"
	log.Println("WebSocket websocket_chat running on ws://localhost:" + port + "/ws")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
