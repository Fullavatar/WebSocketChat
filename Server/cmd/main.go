package main

import (
	"github.com/fullavatar/realTimeChat"
	"log"
	"net/http"
)

func main() {
	server := realTimeChat.NewWebSocketServer()

	http.Handle("/ws", server)

	port := "5000"
	log.Println("WebSocket server running on ws://localhost:" + port + "/ws")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
