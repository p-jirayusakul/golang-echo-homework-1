package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/validator"

	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/config"
	user_handler "github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/handlers"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/handlers/grpc_client"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/repositories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/repositories/factories"
)

func main() {

	// Load config
	configs.LoadConfig(".env")

	// App
	app := echo.New()

	// Repository
	db := config.InitDatabase()
	repoAccount := repositories.NewAccountRepository(db)
	repoResetPassword := repositories.NewResetPasswordRepository(db)

	// Middlewere
	app.Validator = validator.NewCustomValidator()
	app.Use(pkg_middleware.ErrorHandler)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(pkg_middleware.LogHandler(logger))

	grpcClient, errs := grpc_client.NewGrpcClient()
	if len(errs) > 0 {
		log.Fatalf("did not connect grpc client: %v", errs)
		panic(errs)
	}

	grpcFactory := factories.NewGrpcFactory(grpcClient)

	// Handler
	user_handler.NewAuthHttpHandler(app, &repoAccount, &repoResetPassword, grpcFactory)
	app.Logger.Fatal(app.Start(":3001"))
}
