package wsManage

import (
	"log"
	"net/http"
)

var ChatsHub = Hub{
	register:    make(chan *Conn),
	unregister:  make(chan *Conn),
	Connections: make(map[*Conn]bool),
}
var ClientMessagesHub = Hub{
	register:    make(chan *Conn),
	unregister:  make(chan *Conn),
	Connections: make(map[*Conn]bool),
}

func WebsocketInit(url string, logger *log.Logger, GrpcModel *Grpc) {
	GrpcModel.ChatsConnections = &ChatsHub.Connections
	GrpcModel.ClientMessagesConnections = &ClientMessagesHub.Connections
	ChatsHub.Logger = logger
	ClientMessagesHub.Logger = logger

	go ChatsHub.run()
	go ClientMessagesHub.run()

	go GrpcModel.ListenMessageStream()
	http.HandleFunc("/getChats", serveWsGetChats)
	http.HandleFunc("/getMessages", serveWsGetMessages)
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal(err)
	}
}
