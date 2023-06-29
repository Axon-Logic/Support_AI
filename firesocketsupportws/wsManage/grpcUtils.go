package wsManage

import (
	"context"
	"encoding/json"
	"firesocketSupportWs/config"
	pb "firesocketSupportWs/grpcApi"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func (h *Grpc) RecvMessageStream() {
	biStream, err := h.grpcServer.Client.GetMessageStream(context.Background(), &pb.MessageStreamRequest{})
	defer biStream.CloseSend()
	if err == nil {
		for {
			message, err := biStream.Recv()
			if err == nil {
				go h.sendMessagesToConns(message)
			} else {
				break
			}
		}
	}
}

func (h *Grpc) SendMessageToProvider(data pb.ClientMessage) error {
	err := fmt.Errorf("Can't connect to provider")
	if h.grpcServer.Connection.GetState().String() == "READY" {
		_, err = h.grpcServer.Client.PostMessageFromClient(context.Background(), &data)
	}
	return err
}

func (h *Grpc) sendMessagesToConns(message *pb.ServerMessage) {
	Count := false
	if message.MessageCount > 0 {
		Count = true
	}
	ChatJson := ChatJson{Sender: message.Sender, Caption: message.Caption, Message: message.Message, MessageDate: message.MessageDate, MessageCount: message.MessageCount, Count: Count, UserName: message.Sender, FirstName: message.Sender, LastName: message.Sender, Owner: !message.IsClient, MessageType: message.MessageType, MessageId: message.MessageId}

	fmt.Println("Len Conn of Messages", len(*h.ClientMessagesConnections))

	for conn := range *h.ClientMessagesConnections {
		if conn.SubscriptionClient == message.Sender {
			if sendMessage, err := json.Marshal(ChatJson); err == nil {
				conn.Send <- sendMessage
			}
		}
	}

	fmt.Println("Len Conn of Chats", len(*h.ChatsConnections))
	if ChatJson.MessageType != "text" {
		ChatJson.Message = ChatJson.MessageType
	}
	fmt.Println("NewMessage", ChatJson)
	for conn := range *h.ChatsConnections {
		fmt.Println("ChatConn", conn)
		if sendMessage, err := json.Marshal(ChatJson); err == nil {
			conn.Send <- sendMessage
		}
	}
}

func (h *Grpc) ListenMessageStream() {
	for {
		fmt.Println("ListenMessageStream START CONN TO PROVIDER")
		h.checkStateAndSetNewGrpcServer()
		fmt.Println("ListenMessageStream START RECV FROM PROVIDER")
		h.RecvMessageStream()
	}
}

func (h *Grpc) checkStateAndSetNewGrpcServer() {
	for {
		h.SetNewGrpcServer()
		time.Sleep(time.Millisecond * 30)

		if h.grpcServer.Connection.GetState().String() == "READY" {
			fmt.Println("GRPC Connection is Ready")
			break
		}
	}
}

func (h *Grpc) SetNewGrpcServer() {
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
