package websocketserver

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

func (s *WebSocketHandler) handleClient(conn *websocket.Conn) {
	s.addClient(conn)
	go s.handleMessages(conn)
}

func (s *WebSocketHandler) handleMessages(conn *websocket.Conn) {
	_, rawMessage, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}

	var message baseMessage
	err = json.Unmarshal(rawMessage, &message)
	if err != nil {
		log.Println("Failed to read message:", err)
		return
	}
	parsedMessage, err := parseMessage(message.Type, rawMessage, conn)
	if err != nil {
		log.Println("Failed to parse message:", err)
		return
	}

	err = handleMessage(parsedMessage)
	if err != nil {
		log.Println("Failed to handle message:", err)
		return
	}
}

func (s *WebSocketHandler) addClient(conn *websocket.Conn) {

}

func sendMessage(message messageInterface, conn *websocket.Conn) {
	err := conn.WriteJSON(message)
	if err != nil {
		log.Println("Failed to write message:", err)
		return
	}
}
