package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/validator"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/config"
	user_handler "github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/handlers"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/handlers/grpc_server"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories"
)

func main() {

	// Load config
	configs.LoadConfig()

	db := config.InitDatabase()

	// run Grpc Server
	go RunGrpcServer(db)
	// end run Grpc Server

	// App
	app := echo.New()

	// Repository
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

func RunGrpcServer(db *gorm.DB) {

	grpcServer := grpc.NewServer()
	grpc_server.HandlerUserServices(grpcServer, db)

	lis, err := net.Listen("tcp", configs.Config.RPC_USERS_PORT)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	go func() {
		log.Println(fmt.Sprintf("Grpc Server listen to: %s", configs.Config.RPC_USERS_PORT))
		log.Fatal(grpcServer.Serve(lis))
	}()
}
