package websocketserver

type messageRoomLeave baseMessage

func (m messageRoomLeave) getType() messageType {
	return m.Type
}

func handleRoomLeave(message processedMessage) {

}
