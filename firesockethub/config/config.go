package config

import (
	"log"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPC_SERVER_PORT string `env:"GRPC_SERVER_PORT" env-required:"false"`
	GIN_SERVER_PORT  string `env:"GIN_SERVER_PORT" env-required:"false"`
}

var Cfg Config

func LoadConfigTest() {
	err := cleanenv.ReadConfig(".env", &Cfg)
	if err != nil {
		log.Fatalln("LOAD CONFIG TEST LOAD ENV", err)
	}
}

func LoadConfig() {
	Cfg.GRPC_SERVER_PORT = strings.Replace(os.Getenv("GRPC_SERVER_PORT"), `"`, "", -1)
	Cfg.GIN_SERVER_PORT = strings.Replace(os.Getenv("GIN_SERVER_PORT"), `"`, "", -1)
}
