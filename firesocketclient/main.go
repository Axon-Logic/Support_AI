package main

import (
	"firesocketClient/config"
	"firesocketClient/wsManage"
)

func main() {
	// config.LoadConfigTest()
	config.LoadConfig()
	logger := config.LogInto("log/logs.log")
	wsManage.WebsocketInit(config.Cfg.WEBSOCKET_SERVER_URL, &logger)
}
