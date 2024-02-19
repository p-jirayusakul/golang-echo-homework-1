package grpc_server

import (
	userPb "github.com/p-jirayusakul/golang-echo-homework-1/proto/_generated/users"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/usecases/profiles"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func HandlerUserServices(s *grpc.Server, db *gorm.DB) {

	profilesRepo := repositories.NewProfileRepository(db)
	profilesUsecase := profiles.NewProfilesInteractor(&profilesRepo)

	userPb.RegisterUsersServicesServer(s, &server{
		profilesUsecase: profilesUsecase,
	})
}
