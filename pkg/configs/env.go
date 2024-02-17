package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *envConfigs

// struct to map env values
type envConfigs struct {
	DATABASE_PORT     int32  `mapstructure:"DATABASE_PORT"`
	DATABASE_USER     string `mapstructure:"DATABASE_USER"`
	DATABASE_HOST     string `mapstructure:"DATABASE_HOST"`
	DATABASE_PASSWORD string `mapstructure:"DATABASE_PASSWORD"`
	DATABASE_NAME     string `mapstructure:"DATABASE_NAME"`
	JWT_SECRET        string `mapstructure:"JWT_SECRET"`
	SECRET_KEY        string `mapstructure:"SECRET_KEY"`
}

// Call to load the variables from env
func LoadConfig() (err error) {
	// Set the name of the env file
	viper.SetConfigFile(".env")

	// Allow viper to read environment variables
	viper.AutomaticEnv()

	// Read the env file
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading .env file:", err)
		return
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err = viper.Unmarshal(&Config); err != nil {
		fmt.Println("Error reading .env file:", err)
		return
	}

	return
}
