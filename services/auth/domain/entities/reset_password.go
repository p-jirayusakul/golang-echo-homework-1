package entities

import (
	"time"

	"github.com/google/uuid"
)

type ResetPassword struct {
	ResetPasswordID uuid.UUID `json:"resetPasswordId"`
	UserID          uuid.UUID `json:"userId"`
	Email           string    `json:"email"`
	RequestID       uuid.UUID `json:"requestId"`
	ExpiresAt       time.Time `json:"expiresAt"`
	Done            bool      `json:"done"`
}
