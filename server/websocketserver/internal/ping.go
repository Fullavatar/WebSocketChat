package websocketserver

import (
	"encoding/json"
	"fmt"
)

type messagePing messageBase

func (m messagePing) getType() messageType {
	return m.Type
}

func parsePing(jsonData []byte) (messageInterface, error) {
	var message messagePing
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handlePing(message messageInterface) {

}
