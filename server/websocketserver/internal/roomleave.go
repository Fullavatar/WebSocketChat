package websocketserver

import (
	"encoding/json"
	"fmt"
)

type messageRoomLeave processedMessage

func (m messageRoomLeave) getType() messageType {
	return m.Type
}

func parseRoomLeave(jsonData []byte) (messageInterface, error) {
	var message messageRoomLeave
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handleRoomLeave(message processedMessage) {

}
