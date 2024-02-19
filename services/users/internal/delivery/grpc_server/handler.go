package grpc_server

import (
	userPb "github.com/p-jirayusakul/golang-echo-homework-1/proto/_generated/users"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/config"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/factories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/usecases/profiles"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func HandlerUserServices(s *grpc.Server, db *gorm.DB, cfg config.UserConfig) {

	dbFactory := factories.NewDBFactory(db)

	userPb.RegisterUsersServicesServer(s, &server{
		profilesUsecase: profiles.NewProfilesInteractor(&cfg, dbFactory),
	})
}
