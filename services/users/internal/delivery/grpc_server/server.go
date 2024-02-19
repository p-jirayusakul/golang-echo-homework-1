package grpc_server

import (
	usersPb "github.com/p-jirayusakul/golang-echo-homework-1/proto/_generated/users"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/usecases"
)

type server struct {
	usersPb.UnimplementedUsersServicesServer
	profilesUsecase usecases.ProfilesUsecase
}
