package websocketserver

import (
	"github.com/fullavatar/websocketchat/websocketserver/internal"
	"log"
	"net/http"
)

func OpenWebSocketServer(port string) {
	server := websocketserver.NewWebSocketServer()

	http.Handle("/ws", server)

	log.Println("WebSocket server running on ws://localhost:" + port + "/ws")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
