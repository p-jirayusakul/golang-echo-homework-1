package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/validator"

	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/config"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/delivery/grpc_client"
	accounts_handler "github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/delivery/http/delivery/auth"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/repositories/factories"
)

const fileEnv = "env/auth/.env"

var (
	cfg       = config.Auth(fileEnv)
	db        = config.InitDatabase(fileEnv)
	dbFactory = factories.NewDBFactory(db)
)

func main() {

	// run rpc client
	grpcClient, errs := grpc_client.NewGrpcClient(cfg)
	if len(errs) > 0 {
		log.Fatalf("did not connect grpc client: %v", errs)
	}

	// App
	app := echo.New()

	// Middlewere
	app.Validator = validator.NewCustomValidator()
	app.Use(pkg_middleware.ErrorHandler)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(pkg_middleware.LogHandler(logger))

	// setup rounter
	initHandler(app, cfg, grpcClient)

	// Handler
	app.Logger.Fatal(app.Start(cfg.HTTP_PORT))
}

func initHandler(
	app *echo.Echo,
	cfg *config.AuthConfig,
	grpcClient *grpc_client.ServerClient,
) {

	grpcClientFactory := factories.NewGrpcFactory(grpcClient)
	accounts_handler.NewAuthHttpHandler(app, cfg, dbFactory, grpcClientFactory)
}
