package websocketserver

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type messageRoomLeave messageBase

func (m messageRoomLeave) getType() messageType {
	return m.Type
}

func (m messageRoomLeave) getCaller() *websocket.Conn {
	return m.Caller
}

func parseRoomLeave(jsonData []byte) (messageInterface, error) {
	var message messageRoomLeave
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handleRoomLeave(message messageInterface) {

}
