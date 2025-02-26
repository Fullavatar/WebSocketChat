package websocketserver

type messageConnection baseMessage

func (m messageConnection) getType() messageType {
	return m.Type
}

func handleConnection(message processedMessage) {

}
