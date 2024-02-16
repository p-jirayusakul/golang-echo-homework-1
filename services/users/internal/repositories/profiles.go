package repositories

import (
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/models"
	"gorm.io/gorm"
)

type ProfilesRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfilesRepository {
	return ProfilesRepository{db: db}
}

func (x *ProfilesRepository) Create(payload entities.Profiles) error {
	arg := models.Profiles{
		UserID:    payload.UserID,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     payload.Phone,
	}

	result := x.db.Create(&arg)

	return result.Error
}
