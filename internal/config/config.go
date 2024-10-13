package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address      string        `yaml:"address" env-required:"true" env-default:"localhost:8080"`
	Timeout      time.Duration `yaml:"timeout" env-default:"10s"`
	IddleTimeout time.Duration `yaml:"iddle_timeout" env-default:"60s"`
}

func MustLoadConfig() Config {
	confgPath := os.Getenv("CONFIG_PATH")
	if confgPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	//File check
	if _, err := os.Stat(confgPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist: %s", confgPath)
	}

	var conf Config

	if err := cleanenv.ReadConfig(confgPath, &conf); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	return conf
}
