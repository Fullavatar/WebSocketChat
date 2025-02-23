package websocketserver

type messagePrivateChat baseMessage

func (m messagePrivateChat) getType() messageType {
	return m.Type
}

func handlePrivateChat(message processedMessage) {

}
