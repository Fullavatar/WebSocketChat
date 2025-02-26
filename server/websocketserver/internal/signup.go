package websocketserver

type messageSignUp baseMessage

func (m messageSignUp) getType() messageType {
	return m.Type
}

func handleSignUp(message processedMessage) {

}
