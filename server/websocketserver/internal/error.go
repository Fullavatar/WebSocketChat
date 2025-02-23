package websocketserver

type messageError baseMessage

func (m messageError) getType() messageType {
	return m.Type
}

func handleError(message processedMessage) {

}
