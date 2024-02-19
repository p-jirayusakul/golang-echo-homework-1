package config

import (
	"os"

	"github.com/spf13/viper"
)

type AuthConfig struct {
	HTTP_PORT  string `mapstructure:"HTTP_PORT"`
	RPC_USERS  string `mapstructure:"RPC_USERS"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
	SECRET_KEY string `mapstructure:"SECRET_KEY"`
}

func Auth(filename string) *AuthConfig {
	if _, err := os.Stat(filename); err == nil {
		viper.SetConfigFile(filename)
		viper.ReadInConfig()
	} else {
		viper.SetDefault("HTTP_PORT", os.Getenv("HTTP_PORT"))
		viper.SetDefault("RPC_USERS", os.Getenv("RPC_USERS"))
		viper.SetDefault("JWT_SECRET", os.Getenv("JWT_SECRET"))
		viper.SetDefault("SECRET_KEY", os.Getenv("SECRET_KEY"))
	}

	viper.AutomaticEnv()

	var config AuthConfig
	viper.Unmarshal(&config)

	return &config
}
