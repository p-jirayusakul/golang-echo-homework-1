package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/repositories/db/models"
)

func InitDatabase(filename string) *gorm.DB {

	cfg := configs.DatabasePostgres(filename)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Accounts{}, &models.ResetPassword{})

	return db
}
