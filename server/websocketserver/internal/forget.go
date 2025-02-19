package websocketserver

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type messageForget messageBase

func (m messageForget) getType() messageType {
	return m.Type
}

func (m messageForget) getCaller() *websocket.Conn {
	return m.Caller
}

func parseForget(jsonData []byte) (messageInterface, error) {
	var message messageForget
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handleForget(message messageInterface) {

}
