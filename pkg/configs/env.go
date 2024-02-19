package configs

import (
	"fmt"
	"os"

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
func LoadConfig(fileName string) (err error) {

	// fileName := ".env"
	if _, err := os.Stat(fileName); err == nil {
		viper.SetConfigFile(fileName)
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println("Error reading .env file:", err)
			return err
		}
	} else if os.IsNotExist(err) {
		viper.AutomaticEnv()

		// // Define default values (optional)
		// DATABASE_PORT := os.Getenv("DATABASE_PORT")
		// DATABASE_PORT_NUM, err := strconv.Atoi(DATABASE_PORT)
		// if err != nil {
		// 	fmt.Println("Error:", err)
		// 	return err
		// }

		// viper.SetDefault("DATABASE_PORT", int32(DATABASE_PORT_NUM))
		// viper.SetDefault("DATABASE_USER", os.Getenv("DATABASE_USER"))
		// viper.SetDefault("DATABASE_HOST", os.Getenv("DATABASE_HOST"))
		// viper.SetDefault("DATABASE_PASSWORD", os.Getenv("DATABASE_PASSWORD"))
		// viper.SetDefault("DATABASE_NAME", os.Getenv("DATABASE_NAME"))
		// viper.SetDefault("JWT_SECRET", os.Getenv("JWT_SECRET"))
		// viper.SetDefault("SECRET_KEY", os.Getenv("SECRET_KEY"))
		// viper.SetDefault("RPC_USERS", os.Getenv("RPC_USERS"))
		// viper.SetDefault("RPC_USERS", os.Getenv("RPC_USERS"))
		// viper.SetDefault("RPC_USERS_PORT", os.Getenv("RPC_USERS_PORT"))
	}

	if err = viper.Unmarshal(&Config); err != nil {
		fmt.Println("Error reading .env file:", err)
		return
	}

	return
}
