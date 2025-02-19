package websocketserver

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type messagePing messageBase

func (m messagePing) getType() messageType {
	return m.Type
}

func (m messagePing) getCaller() *websocket.Conn {
	return m.Caller
}

func parsePing(jsonData []byte) (messageInterface, error) {
	var message messagePing
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handlePing(message messageInterface) {
	msg, err := buildMessage(PING)
	if err != nil {
		fmt.Println(err)
		return
	}
	sendMessage(msg, message.getCaller())
}
