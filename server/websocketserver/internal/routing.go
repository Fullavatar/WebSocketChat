package websocketserver

import (
	"log"

	"github.com/gorilla/websocket"
)

func (s *WebSocketHandler) handleClient(conn *websocket.Conn) {
	s.addClient(conn)
	go s.handleMessages(conn)
}

func (s *WebSocketHandler) handleMessages(conn *websocket.Conn) {
	var message baseMessage
	err := conn.ReadJSON(message)
	if err != nil {
		log.Println("Failed to read message:", err)
		return
	}
	parsed_message, err := parseMessage(message, conn)
	if err != nil {
		log.Println("Failed to parse message:", err)
		return
	}
	handleMessage(parsed_message)
}

func (s *WebSocketHandler) addClient(conn *websocket.Conn) {

}

func sendMessage(message messageInterface, conn *websocket.Conn) {

}
