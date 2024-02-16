package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/models"
)

func InitDatabase() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok", configs.Config.DATABASE_HOST, configs.Config.DATABASE_USER, configs.Config.DATABASE_PASSWORD, configs.Config.DATABASE_NAME, configs.Config.DATABASE_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Profiles{})

	return db
}
