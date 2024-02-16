package main

import (
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/config"
	profiles_handler "github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/handlers"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories"
)

func main() {

	configs.LoadConfig()

	app := echo.New()

	db := config.InitDatabase()
	repo := repositories.NewProfileRepository(db)
	profiles_handler.NewUserHttpHandler(app, &repo)

	app.Logger.Fatal(app.Start(":3000"))
}
