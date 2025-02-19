package websocketserver

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type messageType string

type messageInterface interface {
	getType() messageType
	getCaller() *websocket.Conn
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

type messageEmpty struct {
	Type messageType `json:"Type"`
}

type messageBase struct {
	Type   messageType `json:"Type"`
	Caller *websocket.Conn
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

func parseMessage(jsonData []byte, conn *websocket.Conn) (messageInterface, error) {
	var message messageEmpty
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse Type: %w", err)
	}
	var message_base messageBase = messageBase{Type: message.Type, Caller: conn}
	if parser, found := messageParsers[message_base.Type]; found {
		return parser(jsonData)
	}
	return nil, fmt.Errorf("unknown message type: %s", message_base.Type)
}

func handleMessage(message messageInterface) error {
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
