package websocketserver

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type messageError messageBase

func (m messageError) getType() messageType {
	return m.Type
}

func (m messageError) getCaller() *websocket.Conn {
	return m.Caller
}

func parseError(jsonData []byte) (messageInterface, error) {
	var message messageError
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handleError(message messageInterface) {

}
