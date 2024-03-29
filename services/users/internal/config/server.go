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
	} else {
		viper.SetDefault("HTTP_PORT", os.Getenv("HTTP_PORT"))
		viper.SetDefault("RPC_PORT", os.Getenv("RPC_PORT"))
		viper.SetDefault("JWT_SECRET", os.Getenv("JWT_SECRET"))
		viper.SetDefault("SECRET_KEY", os.Getenv("SECRET_KEY"))
	}

	viper.AutomaticEnv()

	var config UserConfig
	viper.Unmarshal(&config)

	return &config
}
