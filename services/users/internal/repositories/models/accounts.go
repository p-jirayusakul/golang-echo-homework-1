package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Accounts struct {
	UserID    uuid.UUID `gorm:"primarykey,type:uuid;default:uuid_generate_v4()"`
	Email     string    `gorm:"uniqueIndex:idx_accounts_email,sort:desc"`
	Password  string
	Address   []Address `gorm:"foreignKey:UserID;references:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
