package repositories

import (
	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
)

type AccountsRepository interface {
	Create(payload entities.Accounts) (uuid.UUID, error)
	Read(email string) (entities.Accounts, error)
}
