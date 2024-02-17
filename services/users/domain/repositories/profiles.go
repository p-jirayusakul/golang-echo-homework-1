package repositories

import (
	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
)

type ProfilesRepository interface {
	Create(payload entities.Profiles) error
	Read(id uuid.UUID) (entities.Profiles, error)
	Update(payload entities.UpdateProfilesDTO) error
	Delete(id uuid.UUID) error
}
