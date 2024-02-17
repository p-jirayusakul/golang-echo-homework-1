package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/validator"

	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/config"
	user_handler "github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/handlers"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories"
)

func main() {

	// Load config
	configs.LoadConfig()

	// App
	app := echo.New()

	// Repository
	db := config.InitDatabase()
	repoProfile := repositories.NewProfileRepository(db)
	repoAddress := repositories.NewAddressRepository(db)

	// Middlewere
	app.Validator = validator.NewCustomValidator()
	app.Use(pkg_middleware.ErrorHandler)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(pkg_middleware.LogHandler(logger))

	// Handler
	user_handler.NewUserHttpHandler(app, &repoProfile, &repoAddress)
	app.Logger.Fatal(app.Start(":3002"))
}
