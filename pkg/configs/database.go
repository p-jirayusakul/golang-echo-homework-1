package configs

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type DatabasePostgresConfig struct {
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     int    `mapstructure:"DATABASE_PORT"`
	Database string `mapstructure:"DATABASE_NAME"`
	User     string `mapstructure:"DATABASE_USER"`
	Password string `mapstructure:"DATABASE_PASSWORD"`
}

func ConvertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func DatabasePostgres(filename string) DatabasePostgresConfig {

	if _, err := os.Stat(filename); err == nil {
		viper.SetConfigFile(filename)
		viper.ReadInConfig()
	}

	var config DatabasePostgresConfig
	viper.Unmarshal(&config)

	return config
}
