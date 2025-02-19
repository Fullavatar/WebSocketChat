package websocketserver

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type messagePrivateChat messageBase

func (m messagePrivateChat) getType() messageType {
	return m.Type
}

func (m messagePrivateChat) getCaller() *websocket.Conn {
	return m.Caller
}

func parsePrivateChat(jsonData []byte) (messageInterface, error) {
	var message messagePrivateChat
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handlePrivateChat(message messageInterface) {

}
