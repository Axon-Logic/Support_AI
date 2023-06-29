package wsManage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait   = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
)

var upgraderMessages = websocket.Upgrader{}

var upgraderChats = websocket.Upgrader{}

func serveWsGetChats(w http.ResponseWriter, r *http.Request) {
	upgraderChats.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgraderChats.Upgrade(w, r, nil)

	if err != nil {
		ChatsHub.Logger.Println(err)
		return
	}

	conn := &Conn{
		Send: make(chan []byte, 256),
		WS:   ws,
	}
	fmt.Println("Created New Conn", conn)
	ChatsHub.register <- conn
	fmt.Println("Connected WebSocket to Get Chats")
	go conn.writePump()
	conn.readGetChatsPump()
	fmt.Println("DisConnected WebSocket to Get Chats")
}
func serveWsGetMessages(w http.ResponseWriter, r *http.Request) {
	upgraderMessages.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgraderMessages.Upgrade(w, r, nil)

	if err != nil {
		ClientMessagesHub.Logger.Println(err)
		return
	}

	conn := &Conn{
		Send: make(chan []byte, 256),
		WS:   ws,
	}

	ClientMessagesHub.register <- conn
	fmt.Println("Connected WebSocket to Get Messages")
	go conn.writePump()
	conn.readClientMessagesPump()
	fmt.Println("DisConnected WebSocket to Get Messages")
}

func (conn *Conn) readClientMessagesPump() {

	defer func() {
		ClientMessagesHub.unregister <- conn
		conn.WS.Close()
	}()
	conn.WS.SetReadDeadline(time.Now().Add(pongWait))
	conn.WS.SetPongHandler(func(string) error { conn.WS.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := conn.WS.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				ClientMessagesHub.Logger.Printf("error: %v id:%v", err, conn.Send)
			}
			break
		}
		var postData GetMessageByIdJson
		if err = json.Unmarshal(message, &postData); err == nil && postData.Sender != "" {
			conn.SubscriptionClient = string(postData.Sender)
		}
	}
}

func (conn *Conn) readGetChatsPump() {

	defer func() {
		ClientMessagesHub.unregister <- conn
		conn.WS.Close()
	}()
	conn.WS.SetReadDeadline(time.Now().Add(pongWait))
	conn.WS.SetPongHandler(func(string) error { conn.WS.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		if _, _, err := conn.WS.ReadMessage(); err != nil {
			break
		}
	}
}

func (conn *Conn) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		conn.WS.Close()
	}()

	for {
		select {
		case message, ok := <-conn.Send:
			if !ok {
				// The hub closed the channel.
				conn.WS.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := conn.WS.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}

		case <-ticker.C:
			if err := conn.WS.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
