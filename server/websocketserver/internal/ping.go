package websocketserver

import (
	"fmt"
)

type messagePing struct {
	Type messageType
}

func (m messagePing) getType() messageType {
	return m.Type
}

func handlePing(message processedMessage) {
	msg, err := buildMessage(PING)
	if err != nil {
		fmt.Println(err)
		return
	}
	sendMessage(msg, message.Caller)
}
