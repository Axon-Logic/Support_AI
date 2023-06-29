package wsManage

import "time"

func (h *Hub) PrintConnectionsCount() {
	for range time.Tick(time.Second * 5) {
		hub.Logger.Println("Connections Count", len(hub.connections))
	}
}

func (h *Hub) run() {
	for {
		select {
		case conn := <-h.register:
			h.connections[conn] = true

		case conn := <-h.unregister:
			for group := range h.SubscribeList {
				h.Unsubscribe(group, conn)
			}
			if _, ok := h.connections[conn]; ok {
				delete(h.connections, conn)
				close(conn.Send)
			}
		}
	}
}

func (h *Hub) Subscribe(Group string, conn *Conn) {
	h.Lock()
	if _, ok := h.SubscribeList[Group]; !ok {
		h.SubscribeList[Group] = make(map[*Conn]bool, 256)
	}
	h.SubscribeList[Group][conn] = true
	h.Unlock()
}

func (h *Hub) Unsubscribe(Group string, conn *Conn) {
	h.Lock()
	if _, ok := h.SubscribeList[Group]; ok {
		delete(h.SubscribeList[Group], conn)
	}
	h.Unlock()
}
