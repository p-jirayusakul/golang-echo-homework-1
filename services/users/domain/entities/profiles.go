package entities

import (
	"github.com/google/uuid"
)

type Profiles struct {
	UserID    uuid.UUID `json:"userId"`
	FirstName *string   `json:"firstName"`
	LastName  *string   `json:"lastName"`
	Email     string    `json:"email"`
	Phone     *string   `json:"phone"`
}

type UpdateProfilesDTO struct {
	UserID    uuid.UUID `json:"userId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

type DeleteProfilesDTO struct {
	UserID uuid.UUID `json:"userId"`
}
