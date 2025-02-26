package websocketserver

type messageChat baseMessage

func (m messageChat) getType() messageType {
	return m.Type
}

func handleChat(message processedMessage) {

}
