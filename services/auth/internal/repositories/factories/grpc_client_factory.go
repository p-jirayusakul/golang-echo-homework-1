package factories

import (
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/repositories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/handlers/grpc_client"
	repository_grpc "github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/repositories/grpc"
)

type GrpcClientFactory struct {
	UsersRepo repositories.UsersRepository
}

func NewGrpcFactory(client *grpc_client.ServerClient) *GrpcClientFactory {
	var (
		usersRepoGrpc repository_grpc.UsersRepository
	)

	usersRepoGrpc = repository_grpc.NewUsersRepository(client.UsersClient)
	return &GrpcClientFactory{
		UsersRepo: &usersRepoGrpc,
	}
}
