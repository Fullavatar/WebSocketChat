package websocketserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type messageType string

const (
	SIGNUP        messageType = "SignUp"
	FORGET        messageType = "Forget"
	CONNECTION    messageType = "Connection"
	DISCONNECTION messageType = "Disconnection"
	PING          messageType = "Ping"
	CHAT          messageType = "Chat"
	PRIVATECHAT   messageType = "PrivateChat"
	STATUS        messageType = "Status"
	ERROR         messageType = "Error"
	ROOMJOIN      messageType = "RoomJoin"
	ROOMLEAVE     messageType = "RoomLeave"
)

type messageInterface interface {
	getType() messageType
}

type baseMessage struct {
	Type messageType `json:"Type"`
}

type processedMessage struct {
	Caller  *websocket.Conn
	Message messageInterface
}

var messageStructs = map[messageType]func() messageInterface{
	PING:          func() messageInterface { return &messagePing{} },
	CHAT:          func() messageInterface { return &messageChat{} },
	PRIVATECHAT:   func() messageInterface { return &messagePrivateChat{} },
	STATUS:        func() messageInterface { return &messageStatus{} },
	ERROR:         func() messageInterface { return &messageError{} },
	ROOMJOIN:      func() messageInterface { return &messageRoomJoin{} },
	ROOMLEAVE:     func() messageInterface { return &messageRoomLeave{} },
	SIGNUP:        func() messageInterface { return &messageSignUp{} },
	FORGET:        func() messageInterface { return &messageForget{} },
	CONNECTION:    func() messageInterface { return &messageConnection{} },
	DISCONNECTION: func() messageInterface { return &messageDisconnection{} },
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

func parseMessage(mType messageType, rawMessage []byte, conn *websocket.Conn) (processedMessage, error) {
	if structure, found := messageStructs[mType]; found {
		message := structure()
		err := json.Unmarshal(rawMessage, &message)
		if err != nil {
			return processedMessage{}, fmt.Errorf("failed to parse message: %w", err)
		}
		return processedMessage{Caller: conn, Message: message}, nil
	}
	return processedMessage{}, fmt.Errorf("failed to parse message: unknown message type: %v", mType)
}

func handleMessage(message processedMessage) error {
	if handler, found := messageHandlers[message.Message.getType()]; found {
		handler(message)
		return nil
	}
	return fmt.Errorf("unknown message type: %s", message.Message.getType())
}

func buildMessage(typeToBuild messageType) (messageInterface, error) {
	if builder, found := messageBuilders[typeToBuild]; found {
		return builder(), nil
	}
	return nil, fmt.Errorf("unknown message type: %s", typeToBuild)
}
