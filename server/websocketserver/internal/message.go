package websocketserver

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type messageType string

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

type messageInterface interface {
	getType() messageType
}

type baseMessage struct {
	Type messageType     `json:"Type"`
	Data json.RawMessage `json:"Data"`
}

type processedMessage struct {
	Caller  *websocket.Conn
	Type    messageType
	Message messageInterface
}

func (m processedMessage) getType() messageType {
	return m.Type
}

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

var messageHandlers = map[messageType]func(processedMessage){
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
	PING:          func() messageInterface { return &messagePing{Type: PING} },
	CHAT:          func() messageInterface { return &messageChat{Type: CHAT} },
	PRIVATECHAT:   func() messageInterface { return &messagePrivateChat{Type: PRIVATECHAT} },
	STATUS:        func() messageInterface { return &messageStatus{Type: STATUS} },
	ERROR:         func() messageInterface { return &messageError{Type: ERROR} },
	ROOMJOIN:      func() messageInterface { return &messageRoomJoin{Type: ROOMJOIN} },
	ROOMLEAVE:     func() messageInterface { return &messageRoomLeave{Type: ROOMLEAVE} },
	SIGNUP:        func() messageInterface { return &messageSignUp{Type: SIGNUP} },
	FORGET:        func() messageInterface { return &messageForget{Type: FORGET} },
	CONNECTION:    func() messageInterface { return &messageConnection{Type: CONNECTION} },
	DISCONNECTION: func() messageInterface { return &messageDisconnection{Type: DISCONNECTION} },
}

func parseMessage(message baseMessage, conn *websocket.Conn) (processedMessage, error) {
	if parser, found := messageParsers[message.Type]; found {
		completedMessage, err := parser(message.Data)
		if err != nil {
			return processedMessage{}, fmt.Errorf("failed to parse message: %w", err)
		}
		return processedMessage{Type: message.Type, Caller: conn, Message: completedMessage}, nil
	}

	return processedMessage{}, fmt.Errorf("unknown message type: %s", message.Type)
}

func handleMessage(message processedMessage) error {
	if handler, found := messageHandlers[message.getType()]; found {
		handler(message)
		return nil
	}
	return fmt.Errorf("unknown message type: %s", message.getType())
}

func buildMessage(type_to_build messageType) (messageInterface, error) {
	if builder, found := messageBuilders[type_to_build]; found {
		return builder(), nil
	}
	return nil, fmt.Errorf("unknown message type: %s", type_to_build)
}
