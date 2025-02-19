package websocketserver

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type messageRoomJoin messageBase

func (m messageRoomJoin) getType() messageType {
	return m.Type
}

func (m messageRoomJoin) getCaller() *websocket.Conn {
	return m.Caller
}

func parseRoomJoin(jsonData []byte) (messageInterface, error) {
	var message messageRoomJoin
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handleRoomJoin(message messageInterface) {

}
