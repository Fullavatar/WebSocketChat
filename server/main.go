package main

import websocketserver "github.com/fullavatar/websocketchat/websocketserver/pkg"

func main() {
	websocketserver.OpenWebSocketServer("5000")
}
