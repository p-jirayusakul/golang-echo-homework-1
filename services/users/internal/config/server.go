package config

import (
	"os"

	"github.com/spf13/viper"
)

type UserConfig struct {
	HTTP_PORT  string `mapstructure:"HTTP_PORT"`
	RPC_PORT   string `mapstructure:"RPC_PORT"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
	SECRET_KEY string `mapstructure:"SECRET_KEY"`
}

func User(filename string) *UserConfig {
	if _, err := os.Stat(filename); err == nil {
		viper.SetConfigFile(filename)
		viper.ReadInConfig()
	}

	viper.AutomaticEnv()

	var config UserConfig
	viper.Unmarshal(&config)

	return &config
}
