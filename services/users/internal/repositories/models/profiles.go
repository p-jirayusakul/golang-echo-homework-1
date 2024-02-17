package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profiles struct {
	UserID    uuid.UUID `gorm:"primarykey,type:uuid;default:uuid_generate_v4();uniqueIndex:pk_profiles_id,sort:desc"`
	Address   []Address `gorm:"foreignKey:UserID;references:UserID"`
	FirstName *string
	LastName  *string
	Email     string `gorm:"uniqueIndex:idx_email,sort:desc"`
	Phone     *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
