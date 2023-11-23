package application

import (
	env "github.com/caitlinelfring/go-env-default"
)

type Config struct {
	ServerPort int
}

func LoadConfig() Config {
	cfg := Config{
		ServerPort: env.GetIntDefault("SERVERPORT", 3000),
	}

	return cfg
}
