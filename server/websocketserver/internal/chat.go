package websocketserver

import (
	"encoding/json"
	"fmt"
)

type messageChat processedMessage

func (m messageChat) getType() messageType {
	return m.Type
}

func parseChat(jsonData []byte) (messageInterface, error) {
	var message messageChat
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handleChat(message processedMessage) {

}
