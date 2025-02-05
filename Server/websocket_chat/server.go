package websocket_chat

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

type WebSocketHandler struct {
	upgrader websocket.Upgrader
	clients  map[*websocket.Conn]string
	mutex    sync.Mutex
}

func NewWebSocketServer() *WebSocketHandler {
	return &WebSocketHandler{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		clients: make(map[*websocket.Conn]string),
	}
}

func (s *WebSocketHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/ws" {
		http.Error(writer, "Invalid WebSocket path", http.StatusNotFound)
		return
	}

	conn, err := s.upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println("Failed to upgrade:", err)
		return
	}
	defer s.removeClient(conn)

	_, nickname, err := conn.ReadMessage()
	if err != nil {
		log.Println("Failed to read nickname:", err)
		return
	}

	if string(nickname) == "" {
		return
	}

	s.mutex.Lock()
	s.clients[conn] = string(nickname)
	s.mutex.Unlock()

	log.Println(Yellow, string(nickname), Reset, "joined the chat")

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(Yellow, s.clients[conn], Reset, "disconnected")
			break
		}
		s.broadcastMessage(s.clients[conn], message)
	}
}

func (s *WebSocketHandler) removeClient(conn *websocket.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	nickname := s.clients[conn]
	if nickname == "" {
		nickname = "Unknown User"
	}
	delete(s.clients, conn)
	_ = conn.Close()
	log.Println(Yellow, nickname, Reset, "left the chat")
}

func (s *WebSocketHandler) broadcastMessage(sender string, message []byte) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	fullMessage := sender + ": " + string(message)

	for client := range s.clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(fullMessage))
		if err != nil {
			log.Println(Red+"Broadcast error:"+Reset, err)
			_ = client.Close()
			delete(s.clients, client)
		}
	}
}
