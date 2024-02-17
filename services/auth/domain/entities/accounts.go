package entities

import "github.com/google/uuid"

type Accounts struct {
	UserID   uuid.UUID `json:"userId"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type UpdatePasswordAccountDTO struct {
	UserID    uuid.UUID `json:"userId"`
	RequestID string    `json:"requestId"`
	Password  string    `json:"password"`
}
