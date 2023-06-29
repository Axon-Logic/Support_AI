package grpcMessage

import (
	"context"
	pb "firesocketHub/grpcApi"
	"fmt"
	"io"
)

func (s *Server) RegisterProvider(stream pb.Main_RegisterProviderServer) error {
	provider, err := stream.Recv()
	if err == nil {
		fmt.Println("RegisterProvider", provider.HostName, provider.GrpcPort)
		if provider.HostName != "" && provider.GrpcPort != "" {
			s.AddProvider(provider)
			defer s.RemoveProvider(provider)
			for {
				err = stream.RecvMsg(nil)
				if err == io.EOF {
					s.Logger.Printf("Stream CONNECTION REFUSED - %v\n", err)
					break
				} else if err != nil {
					s.Logger.Printf("Stream ERROR - %v\n", err)
					break
				}
				fmt.Println("Stream OK", err)
			}
		}
	}

	return stream.Context().Err()
}
func (s *Server) GetMasterProvider(context.Context, *pb.Empty) (*pb.Provider, error) {
	return s.GetProvider(), nil
}
