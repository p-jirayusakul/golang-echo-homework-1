package config

import (
	"os"
)

type UserConfig struct {
	HttpPort string
	RpcPort  string
}

func User() *UserConfig {
	return &UserConfig{
		HttpPort: os.Getenv("HTTP_PORT"),
		RpcPort:  os.Getenv("RPC_PORT"),
	}
}
