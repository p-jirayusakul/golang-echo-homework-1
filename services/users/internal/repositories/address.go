package repositories

import (
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/models"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return AddressRepository{db: db}
}

func (x *AddressRepository) Create(payload entities.Address) error {
	arg := models.Address{
		UserID:   payload.UserID,
		AddrType: payload.AddrType,
		AddrNo:   payload.AddrNo,
		Street:   payload.Street,
		City:     payload.City,
		State:    payload.State,
	}

	result := x.db.Create(&arg)

	return result.Error
}
