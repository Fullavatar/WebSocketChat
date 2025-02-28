package websocketserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fullavatar/websocketchat/database"
	models "github.com/fullavatar/websocketchat/database/pkg"
	websocketserver "github.com/fullavatar/websocketchat/websocketserver/internal"
)

func OpenWebSocketServer(port string) {
	startDataBase()
	server := websocketserver.NewWebSocketServer()

	http.Handle("/ws", server)

	log.Println("WebSocket server running on ws://localhost:" + port + "/ws")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server error:", err)
	}
}

func startDataBase() {
	fmt.Println("Connecting to database...")
	database.Connect()

	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database migrated")
}
