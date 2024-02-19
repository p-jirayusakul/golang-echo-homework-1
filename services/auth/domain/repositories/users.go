package repositories

import (
	"context"

	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
)

type UsersRepository interface {
	CreateProfiles(ctx context.Context, payload *entities.Profiles) error
}
