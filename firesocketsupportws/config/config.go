package config

import (
	"log"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	WEBSOCKET_SERVER_URL string `env:"WEBSOCKET_SERVER_URL" env-required:"false"`
	HUB_GRPC_URL         string `env:"HUB_GRPC_URL" env-required:"false"`
	GIN_SERVER_PORT      string `env:"GIN_SERVER_PORT" env-required:"false"`
	HISTORY_SERVICE_URL  string `env:"HISTORY_SERVICE_URL" env-required:"false"`
}

var Cfg Config

func LoadConfigTest() {
	err := cleanenv.ReadConfig(".env", &Cfg)
	if err != nil {
		log.Fatalln("LOAD CONFIG TEST LOAD ENV", err)
	}
}

func LoadConfig() {
	Cfg.HUB_GRPC_URL = strings.Replace(os.Getenv("HUB_GRPC_URL"), `"`, "", -1)
	Cfg.WEBSOCKET_SERVER_URL = strings.Replace(os.Getenv("WEBSOCKET_SERVER_URL"), `"`, "", -1)
	Cfg.GIN_SERVER_PORT = strings.Replace(os.Getenv("GIN_SERVER_PORT"), `"`, "", -1)
	Cfg.HISTORY_SERVICE_URL = strings.Replace(os.Getenv("HISTORY_SERVICE_URL"), `"`, "", -1)
}
