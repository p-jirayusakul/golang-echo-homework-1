package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/labstack/echo/v4"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/validator"
	"google.golang.org/grpc"

	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/config"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/delivery/grpc_server"
	address_handler "github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/delivery/http/delivery/address"
	profiles_handler "github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/delivery/http/delivery/profiles"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/factories"
)

const fileEnv = "env/users/.env"

var (
	cfg       = config.User(fileEnv)
	db        = config.InitDatabase(fileEnv)
	dbFactory = factories.NewDBFactory(db)
)

func main() {

	// run Grpc Server
	go RunGrpcServer()
	// end run Grpc Server

	// App
	app := echo.New()

	// Middlewere
	app.Validator = validator.NewCustomValidator()
	app.Use(pkg_middleware.ErrorHandler)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(pkg_middleware.LogHandler(logger))

	// setup rounter
	initHandler(app, cfg)

	// Handler
	app.Logger.Fatal(app.Start(":" + cfg.HTTP_PORT))
}

func RunGrpcServer() {

	grpcServer := grpc.NewServer()
	grpc_server.HandlerUserServices(grpcServer, db, *cfg)

	lis, err := net.Listen("tcp", cfg.RPC_PORT)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	go func() {
		log.Println(fmt.Sprintf("Grpc Server listen to: %s", cfg.RPC_PORT))
		log.Fatal(grpcServer.Serve(lis))
	}()
}

func initHandler(
	app *echo.Echo,
	cfg *config.UserConfig) {

	profiles_handler.NewProfilesHttpHandler(app, cfg, dbFactory)
	address_handler.NewAddressHttpHandler(app, cfg, dbFactory)
}
