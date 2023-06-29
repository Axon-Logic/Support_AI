package wsManage

import (
	pb "firesocketSupportWs/grpcApi"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

type Grpc struct {
	ChatsConnections          *map[*Conn]bool
	ClientMessagesConnections *map[*Conn]bool
	grpcServer                GrpcServer
}

type Hub struct {
	// Registered connections.
	Connections map[*Conn]bool

	// Register requests from the connections.
	register chan *Conn

	// Unregister requests from connections.
	unregister chan *Conn
	Logger     *log.Logger
	sync.Mutex
}

type Conn struct {
	// The websocket connection.
	WS *websocket.Conn

	// Buffered channel of outbound messages.
	Send               chan []byte
	SubscriptionClient string
}

type GrpcServer struct {
	Client     pb.MainClient
	Connection *grpc.ClientConn
}