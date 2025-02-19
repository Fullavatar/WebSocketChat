package websocketserver

import (
	"encoding/json"
	"fmt"
)

type messageType string

type messageInterface interface {
	getType() messageType
}

type messageBase struct {
	Type messageType `json:"Type"`
}

const (
	SIGNUP        messageType = "signup"
	FORGET        messageType = "forget"
	CONNECTION    messageType = "connection"
	DISCONNECTION messageType = "disconnection"
	PING          messageType = "ping"
	CHAT          messageType = "chat"
	PRIVATECHAT   messageType = "privatechat"
	STATUS        messageType = "status"
	ERROR         messageType = "error"
	ROOMJOIN      messageType = "roomjoin"
	ROOMLEAVE     messageType = "roomleave"
)

var messageParsers = map[messageType]func([]byte) (messageInterface, error){
	PING:          parsePing,
	CHAT:          parseChat,
	PRIVATECHAT:   parsePrivateChat,
	STATUS:        parseStatus,
	ERROR:         parseError,
	ROOMJOIN:      parseRoomJoin,
	ROOMLEAVE:     parseRoomLeave,
	SIGNUP:        parseSignUp,
	FORGET:        parseForget,
	CONNECTION:    parseConnection,
	DISCONNECTION: parseDisconnection,
}

var messageHandlers = map[messageType]func(messageInterface){
	PING:          handlePing,
	CHAT:          handleChat,
	PRIVATECHAT:   handlePrivateChat,
	STATUS:        handleStatus,
	ERROR:         handleError,
	ROOMJOIN:      handleRoomJoin,
	ROOMLEAVE:     handleRoomLeave,
	SIGNUP:        handleSignUp,
	FORGET:        handleForget,
	CONNECTION:    handleConnection,
	DISCONNECTION: handleDisconnection,
}

var messageBuilders = map[messageType]func() messageInterface{
	PING:          buildPing,
	CHAT:          buildChat,
	PRIVATECHAT:   buildPrivateChat,
	STATUS:        buildStatus,
	ERROR:         buildError,
	ROOMJOIN:      buildRoomJoin,
	ROOMLEAVE:     buildRoomLeave,
	SIGNUP:        buildSignUp,
	FORGET:        buildForget,
	CONNECTION:    buildConnection,
	DISCONNECTION: buildDisconnection,
}

func parseMessage(jsonData []byte) (messageInterface, error) {
	var message messageBase
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse Type: %w", err)
	}

	if perser, found := messageParsers[message.Type]; found {
		return perser(jsonData)
	}

	return nil, fmt.Errorf("unknown message type: %s", message.Type)
}

func handleMessage(message messageInterface) error {
	if handler, found := messageHandlers[message.getType()]; found {
		handler(message)
		return nil
	}

	return fmt.Errorf("unknown message type: %s", message.getType())
}

func buildMessage(type_to_build messageType) {

}
