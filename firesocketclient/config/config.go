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
}

var Cfg Config

func LoadConfigTest() {
	err := cleanenv.ReadConfig(".env", &Cfg)
	if err != nil {
		log.Fatalln("LOAD CONFIG TEST LOAD ENV", err)
	}
}

func LoadConfig() {
	Cfg.WEBSOCKET_SERVER_URL = strings.Replace(os.Getenv("WEBSOCKET_SERVER_URL"), `"`, "", -1)
	Cfg.HUB_GRPC_URL = strings.Replace(os.Getenv("HUB_GRPC_URL"), `"`, "", -1)
}
