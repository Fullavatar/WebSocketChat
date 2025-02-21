package websocketserver

import (
	"encoding/json"
	"fmt"
)

type messagePing struct {
	Type messageType
}

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

func handlePing(message processedMessage) {
	msg, err := buildMessage(PING)
	if err != nil {
		fmt.Println(err)
		return
	}
	sendMessage(msg, message.Caller)
}
