package websocketserver

type messageDisconnection baseMessage

func (m messageDisconnection) getType() messageType {
	return m.Type
}

func handleDisconnection(message processedMessage) {

}
