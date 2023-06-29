package wsManage

import "time"

func (h *Hub) PrintConnectionsCount() {
	for range time.Tick(time.Second * 5) {
		h.Logger.Println("Connections Count", len(h.Connections))
	}
}

func (h *Hub) run() {
	for {
		select {
		case conn := <-h.register:
			h.Lock()
			h.Connections[conn] = true
			h.Unlock()

		case conn := <-h.unregister:
			if _, ok := h.Connections[conn]; ok {
				delete(h.Connections, conn)
				close(conn.Send)
			}
		}
	}
}
