package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port int    `env:"PORT" env-default:"8080"`
	Host string `env:"HOST" env-default:"localhost"`
}

func main() {
	var cfg Config

	// Load config from environment variables
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	fmt.Printf("Config Loaded: %s:%d\n", cfg.Host, cfg.Port)
}
