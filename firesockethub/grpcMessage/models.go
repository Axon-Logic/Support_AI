package grpcMessage

import (
	pb "firesocketHub/grpcApi"
	"fmt"
	"log"
	"sync"
)

type Server struct {
	pb.UnimplementedMainServer
	Providers []*pb.Provider
	Logger    log.Logger
	sync.Mutex
}

func (s *Server) AddProvider(provider *pb.Provider) int {
	s.Lock()
	index := len(s.Providers)
	s.Providers = append(s.Providers, provider)
	s.Unlock()
	return index
}

func RemoveIndex(s []*pb.Provider, index int) []*pb.Provider {
	return append(s[:index], s[index+1:]...)
}

func (s *Server) RemoveProvider(provider *pb.Provider) {
	s.Lock()
	for i, p := range s.Providers {
		if p.HostName == provider.HostName && p.GrpcPort == provider.GrpcPort {
			fmt.Println("RemoveProvider", p.GrpcPort, i)
			s.Providers = RemoveIndex(s.Providers, i)
			break
		}
	}
	s.Unlock()
}

func (s *Server) GetProvider() *pb.Provider {
	s.Lock()
	provider := &pb.Provider{}
	if len(s.Providers) > 0 {
		provider = s.Providers[0]
	}
	s.Unlock()
	return provider
}
