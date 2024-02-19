package configs

import (
	"os"
	"strconv"
)

type DatabasePostgresConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func ConvertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func DatabasePostgres() DatabasePostgresConfig {
	return DatabasePostgresConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     ConvertInt(os.Getenv("DATABASE_PORT")),
		Database: os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	}
}
