package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}

type Config struct {
	Env          string     `yaml:"env" env:"ENV" env-default:"production" env-required:"true"`
	Storage_path string     `yaml:"storage_path" env-required:"true"`
	HTTPServer   HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == ""{
		flags := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == ""{
			log.Fatal("config path is required")

		}
	}

	if _,err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("config file does not exist: %s",configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil{
		log.Fatalf("failed to read config: %v",err)
	}
	// ab agar sab kuch sahi chal raha hai to config return kardo

	return &cfg
}