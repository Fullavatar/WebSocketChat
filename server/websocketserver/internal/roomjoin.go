package websocketserver

type messageRoomJoin baseMessage

func (m messageRoomJoin) getType() messageType {
	return m.Type
}

func handleRoomJoin(message processedMessage) {

}
