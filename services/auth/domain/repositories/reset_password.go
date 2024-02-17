package repositories

import (
	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
)

type ResetPasswordRepository interface {
	Create(payload entities.ResetPassword) (uuid.UUID, error)
	Read(requestId uuid.UUID) (entities.ResetPassword, error)
	UpdateDone(requestId uuid.UUID) error
}
