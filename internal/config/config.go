package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Storage string `yaml:"storage_path"`
	Env     string `yaml:"env"`
}

func InitConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		fmt.Printf("CONFIG_PATH is empty \n")
		os.Exit(1)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Printf("File doest exist %v \n", configPath)
		os.Exit(1)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		fmt.Printf("Cannot read config %v \n", configPath)
		os.Exit(1)
	}

	return &cfg
}
