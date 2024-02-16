package entities

import "github.com/google/uuid"

type Accounts struct {
	UserID   uuid.UUID `json:"userId"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
