package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Accounts struct {
	UserID        uuid.UUID       `gorm:"primarykey,type:uuid;default:uuid_generate_v4();uniqueIndex:pk_accounts_user_id,sort:desc"`
	ResetPassword []ResetPassword `gorm:"foreignKey:UserID;references:UserID"`
	Email         string          `gorm:"uniqueIndex:idx_accounts_email,sort:desc"`
	Password      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
