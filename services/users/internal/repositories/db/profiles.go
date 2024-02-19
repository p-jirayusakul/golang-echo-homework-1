package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/common"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/db/models"
	"gorm.io/gorm"
)

type ProfilesRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfilesRepository {
	return ProfilesRepository{db: db}
}

func (x *ProfilesRepository) Create(payload entities.Profiles) error {
	data := models.Profiles{
		UserID:    payload.UserID,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     payload.Phone,
	}

	return x.db.Create(&data).Error
}

func (x *ProfilesRepository) Read(id uuid.UUID) (entities.Profiles, error) {
	data := models.Profiles{}

	result := x.db.Where("user_id = ?", id).First(&data)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return entities.Profiles{}, common.ErrDataNotFound
	}

	profile := entities.Profiles{
		UserID:    data.UserID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Phone:     data.Phone,
	}

	return profile, nil
}

func (x *ProfilesRepository) Update(payload entities.UpdateProfilesDTO) error {
	data := models.Profiles{}

	result := x.db.Where("user_id = ?", payload.UserID).First(&data)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return common.ErrDataNotFound
	}

	data.FirstName = &payload.FirstName
	data.LastName = &payload.LastName

	return x.db.Model(models.Profiles{}).Where("user_id = ?", payload.UserID.String()).Updates(data).Error
}

func (x *ProfilesRepository) Delete(id uuid.UUID) error {
	return x.db.Where("user_id = ?", id.String()).Delete(&models.Profiles{}).Error
}
