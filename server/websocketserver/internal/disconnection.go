package websocketserver

import (
	"encoding/json"
	"fmt"
)

type messageDisconnection messageBase

func (m messageDisconnection) getType() messageType {
	return m.Type
}

func parseDisconnection(jsonData []byte) (messageInterface, error) {
	var message messageDisconnection
	if err := json.Unmarshal(jsonData, &message); err != nil {
		return nil, fmt.Errorf("failed to parse ping: %w", err)
	}
	return message, nil
}

func handleDisconnection(message messageInterface) {

}
