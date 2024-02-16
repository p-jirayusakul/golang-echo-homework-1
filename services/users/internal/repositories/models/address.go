package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	AddressId uuid.UUID `gorm:"primarykey,type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID
	AddrType  string
	AddrNo    string
	Street    string
	City      string
	State     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
