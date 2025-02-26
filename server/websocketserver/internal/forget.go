package websocketserver

type messageForget baseMessage

func (m messageForget) getType() messageType {
	return m.Type
}

func handleForget(message processedMessage) {

}
