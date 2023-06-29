package main

import (
	"firesocketSupportWs/config"
	ginReg "firesocketSupportWs/ginRegulator"
	"firesocketSupportWs/wsManage"
)

func main() {
	// config.LoadConfigTest()
	config.LoadConfig()
	logger := config.LogInto("log/logs.log")
	GrpcModel := wsManage.Grpc{}

	go ginReg.GinInit(logger, &GrpcModel)
	wsManage.WebsocketInit(config.Cfg.WEBSOCKET_SERVER_URL, &logger, &GrpcModel)
}
