package websocketserver

import (
	"encoding/json"
	"fmt"
)

type messageSignUp messageBase

func (m messageSignUp) getType() messageType {
	return m.Type
}

func parseSignUp(jsonData []byte) (messageInterface, error) {
	var message messageSignUp
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handleSignUp(message messageInterface) {

}
