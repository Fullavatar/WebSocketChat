package websocketserver

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type messageStatus messageBase

func (m messageStatus) getType() messageType {
	return m.Type
}

func (m messageStatus) getCaller() *websocket.Conn {
	return m.Caller
}

func parseStatus(jsonData []byte) (messageInterface, error) {
	var message messageStatus
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handleStatus(message messageInterface) {

}
