package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/common"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/db/models"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return AddressRepository{db: db}
}

func (x *AddressRepository) Create(payload entities.Address) error {
	data := models.Address{
		UserID:   payload.UserID,
		AddrType: payload.AddrType,
		AddrNo:   payload.AddrNo,
		Street:   payload.Street,
		City:     payload.City,
		State:    payload.State,
	}

	result := x.db.Create(&data)

	return result.Error
}

func (x *AddressRepository) Read(id uuid.UUID) (entities.Address, error) {
	data := models.Address{}

	result := x.db.Where("address_id = ?", id).First(&data)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return entities.Address{}, common.ErrDataNotFound
	}

	adress := entities.Address{
		AddressId: data.AddressId,
		UserID:    data.UserID,
		AddrType:  data.AddrType,
		AddrNo:    data.AddrNo,
		Street:    data.Street,
		City:      data.City,
		State:     data.State,
	}

	return adress, nil
}

func (x *AddressRepository) Update(arg entities.Address) error {
	data := models.Address{}

	result := x.db.Where("address_id = ?", arg.AddressId).First(&data)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return common.ErrDataNotFound
	}

	data.AddressId = arg.AddressId
	data.AddrNo = arg.AddrNo
	data.AddrType = arg.AddrType
	data.Street = arg.Street
	data.City = arg.City
	data.State = arg.State

	return x.db.Model(models.Address{}).Where("address_id = ?", data.AddressId).Updates(data).Error
}

func (x *AddressRepository) Delete(id uuid.UUID) error {
	return x.db.Where("address_id = ?", id).Delete(&models.Address{}).Error
}
