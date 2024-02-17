package repositories

import (
	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
)

type AddressRepository interface {
	Create(payload entities.Address) error
	Read(id uuid.UUID) (entities.Address, error)
	Update(arg entities.Address) error
	Delete(id uuid.UUID) error
}
