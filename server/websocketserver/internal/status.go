package websocketserver

type messageStatus baseMessage

func (m messageStatus) getType() messageType {
	return m.Type
}

func handleStatus(message processedMessage) {

}
