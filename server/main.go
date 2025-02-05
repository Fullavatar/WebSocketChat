package main

import (
	"log"
	"net/http"

	"github.com/fullavatar/websocketchat/websocketserver"
)

func main() {
	server := websocketserver.NewWebSocketServer()

	http.Handle("/ws", server)

	port := "5000"
	log.Println("WebSocket server running on ws://localhost:" + port + "/ws")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
