package wsManage

import (
	"encoding/json"
	pb "firesocketClient/grpcApi"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait   = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
)

var upgrader = websocket.Upgrader{}

func serveWs(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		hub.Logger.Println(err)
		return
	}

	conn := &Conn{
		Send: make(chan pb.ServerMessage, 256),
		WS:   ws,
	}

	hub.register <- conn
	go conn.writePump()
	conn.readPump()
}

func (conn *Conn) readPump() {
	var (
		recievedData WSRecievedData
		postData     pb.ClientMessage
	)

	defer func() {
		hub.unregister <- conn
		conn.WS.Close()
	}()
	conn.WS.SetReadDeadline(time.Now().Add(pongWait))
	conn.WS.SetPongHandler(func(string) error { conn.WS.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		recievedData = WSRecievedData{}
		postData = pb.ClientMessage{}
		_, message, err := conn.WS.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				hub.Logger.Printf("error: %v id:%v", err, conn.Send)
			}
			break
		}

		if err = json.Unmarshal(message, &postData); err == nil && postData.Sender != "" && postData.Message != "" {
			hub.grpcServer.PostMessage(postData)
			continue
		}

		if err = json.Unmarshal(message, &recievedData); err == nil || recievedData.Sender != "" {
			if recievedData.Subscribe == 1 {
				hub.Subscribe(recievedData.Sender, conn)
			} else if recievedData.Subscribe == 0 {
				hub.Unsubscribe(recievedData.Sender, conn)
			}
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
			if jsonMessage, err := json.Marshal(message); err == nil {
				if err := conn.WS.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
					return
				}
			}

		case <-ticker.C:
			if err := conn.WS.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
