package wsManage

import (
	pb "firesocketClient/grpcApi"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

type Hub struct {
	// Registered connections.
	connections map[*Conn]bool

	// Register requests from the connections.
	register chan *Conn

	// Unregister requests from connections.
	unregister    chan *Conn
	SubscribeList map[string]map[*Conn]bool
	Logger        *log.Logger
	grpcServer    GrpcServer
	sync.Mutex
}

type Conn struct {
	// The websocket connection.
	WS *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan pb.ServerMessage
}

type GrpcServer struct {
	Client     pb.MainClient
	Connection *grpc.ClientConn
}
