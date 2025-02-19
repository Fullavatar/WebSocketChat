package websocketserver

import (
	"encoding/json"
	"fmt"
)

type messageConnection messageBase

func (m messageConnection) getType() messageType {
	return m.Type
}

func parseConnection(jsonData []byte) (messageInterface, error) {
	var message messageConnection
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handleConnection(message messageInterface) {

}
