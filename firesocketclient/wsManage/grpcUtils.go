package wsManage

import (
	"context"
	"firesocketClient/config"
	pb "firesocketClient/grpcApi"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func (s *GrpcServer) PostMessage(Data pb.ClientMessage) {
	if s.Connection.GetState().String() == "READY" {
		s.Client.PostMessageFromClient(context.Background(), &Data)
	}
}

func (h *Hub) RecvMessageStream() {
	biStream, err := h.grpcServer.Client.GetMessageStream(context.Background(), &pb.MessageStreamRequest{})
	defer biStream.CloseSend()
	if err == nil {
		for {
			message, err := biStream.Recv()
			if err == nil {
				connections, ok := h.SubscribeList[message.Sender]
				if ok {
					for conn := range connections {
						conn.Send <- *message
					}
				}
			} else {
				break
			}
		}
	}
}

func (h *Hub) ListenMessageStream() {
	for {
		fmt.Println("Try to Listen GRPC Stream(not Started)...")
		h.checkStateAndSetNewGrpcServer()
		fmt.Println("Start Listening GRPC Stream")
		h.RecvMessageStream()
	}
}

func (h *Hub) checkStateAndSetNewGrpcServer() {
	for {
		h.SetNewGrpcServer()
		time.Sleep(time.Millisecond * 30)

		if h.grpcServer.Connection.GetState().String() == "READY" {
			fmt.Println("GRPC Connection is Ready")
			break
		}
	}
}

func (h *Hub) SetNewGrpcServer() {
	h.grpcServer.Connection, _ = grpc.Dial("0:0", []grpc.DialOption{grpc.WithInsecure()}...)
	provider, err := GetMasterProvider()
	if err == nil {
		if h.grpcServer.Connection, err = grpc.Dial(provider.HostName+":"+provider.GrpcPort, []grpc.DialOption{grpc.WithInsecure()}...); err == nil {
			h.grpcServer.Client = pb.NewMainClient(h.grpcServer.Connection)
		}
	}
}

func GetMasterProvider() (*pb.Provider, error) {
	var provider *pb.Provider
	err := error(nil)
	if grpcClient, err := GetNewGrpcCient(config.Cfg.HUB_GRPC_URL); err == nil {
		if provider, err = grpcClient.GetMasterProvider(context.Background(), &pb.Empty{}); err != nil {
			provider = &pb.Provider{}
		}
	}
	if provider.HostName == "" && provider.GrpcPort == "" {
		err = fmt.Errorf("Can't connect to provider")
	}
	return provider, err
}

func GetNewGrpcCient(url string) (pb.MainClient, error) {
	conn, err := grpc.Dial(url, []grpc.DialOption{grpc.WithInsecure()}...)
	return pb.NewMainClient(conn), err
}
