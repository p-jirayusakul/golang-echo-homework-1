package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/common"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/repositories/models"
	"gorm.io/gorm"
)

type ResetPasswordRepository struct {
	db *gorm.DB
}

func NewResetPasswordRepository(db *gorm.DB) ResetPasswordRepository {
	return ResetPasswordRepository{db: db}
}

func (x *ResetPasswordRepository) Create(payload entities.ResetPassword) (uuid.UUID, error) {
	data := models.ResetPassword{
		UserID:    payload.UserID,
		ExpiresAt: payload.ExpiresAt,
		RequestID: payload.RequestID,
	}

	result := x.db.Create(&data)

	return data.RequestID, result.Error
}

func (x *ResetPasswordRepository) Read(requestId uuid.UUID) (entities.ResetPassword, error) {
	data := models.ResetPassword{}

	result := x.db.Where("request_id = ?", requestId).First(&data)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return entities.ResetPassword{}, common.ErrDataNotFound
	}

	resetpass := entities.ResetPassword{
		ResetPasswordID: data.ResetPasswordID,
		UserID:          data.UserID,
		RequestID:       data.RequestID,
		ExpiresAt:       data.ExpiresAt,
		Done:            data.Done,
	}

	return resetpass, result.Error
}

func (x *ResetPasswordRepository) UpdateDone(requestId uuid.UUID) error {
	data := models.ResetPassword{}

	result := x.db.Where("request_id = ?", requestId).First(&data)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return common.ErrDataNotFound
	}

	data.Done = true

	return x.db.Model(models.ResetPassword{}).Where("request_id = ?", requestId).Updates(data).Error
}
