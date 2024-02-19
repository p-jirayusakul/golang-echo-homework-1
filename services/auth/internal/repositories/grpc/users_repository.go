package grpc

import (
	"context"

	usersPb "github.com/p-jirayusakul/golang-echo-homework-1/proto/_generated/users"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
	"google.golang.org/grpc"
)

type UsersRepository struct {
	svc usersPb.UsersServicesClient
	ctx context.Context
}

func NewUsersRepository(con *grpc.ClientConn) UsersRepository {
	client := usersPb.NewUsersServicesClient(con)

	return UsersRepository{
		svc: client,
		ctx: nil,
	}
}

func (x *UsersRepository) CreateProfiles(ctx context.Context, payload *entities.Profiles) error {
	_, err := x.svc.CreateProfiles(ctx, &usersPb.CreateProfilesRequest{
		UserId: payload.UserID,
		Email:  payload.Email,
	})

	return err
}
