package realTimeChat

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestWebSocketBroadcast(t *testing.T) {
	wsServer := NewWebSocketServer()
	server := httptest.NewServer(wsServer)
	defer server.Close()

	wsURL := strings.Replace(server.URL, "http", "ws", 1) + "/ws"

	client1, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer func(client1 *websocket.Conn) {
		err = client1.Close()
		if err != nil {
			panic(err)
		}
	}(client1)

	client2, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer func(client2 *websocket.Conn) {
		err = client2.Close()
		if err != nil {
			panic(err)
		}
	}(client2)

	client3, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer func(client3 *websocket.Conn) {
		err = client3.Close()
		if err != nil {
			panic(err)
		}
	}(client3)

	testMessage := "Hello, everyone!"
	err = client1.WriteMessage(websocket.TextMessage, []byte(testMessage))
	require.NoError(t, err)

	readMessageWithTimeout := func(conn *websocket.Conn) (string, error) {
		err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		if err != nil {
			return "", err
		}
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return "", err
		}
		return string(msg), nil
	}

	clients := []*websocket.Conn{client2, client3}
	for _, client := range clients {
		receivedMessage, err := readMessageWithTimeout(client)
		require.NoError(t, err, "Client did not receive a message in time")
		assert.Equal(t, testMessage, receivedMessage)
	}
}
