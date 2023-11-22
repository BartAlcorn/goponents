package application

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort uint16
}

func LoadConfig() Config {
	cfg := Config{
		ServerPort: 3000,
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	return cfg
}
