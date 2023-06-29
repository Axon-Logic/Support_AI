package main

import (
	"firesocketHub/config"
	ginReg "firesocketHub/ginRegulator"
	pb "firesocketHub/grpcApi"
	"firesocketHub/grpcMessage"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	config.LoadConfig()
	// config.LoadConfigTest()
	logger := config.LogInto("log/logs.log")
	s := grpc.NewServer()

	StreamServer := &grpcMessage.Server{Logger: logger, Providers: make([]*pb.Provider, 0)}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Cfg.GRPC_SERVER_PORT))
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	go ginReg.GinInit(logger, StreamServer)
	pb.RegisterMainServer(s, StreamServer)
	fmt.Println("GRPC SERVER STARTED", fmt.Sprintf(":%s", config.Cfg.GRPC_SERVER_PORT))
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
