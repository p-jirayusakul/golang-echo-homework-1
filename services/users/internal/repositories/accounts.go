package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/common"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/models"
	"gorm.io/gorm"
)

type AccountsRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountsRepository {
	return AccountsRepository{db: db}
}

func (x *AccountsRepository) Create(payload entities.Accounts) (uuid.UUID, error) {
	data := models.Accounts{
		Email:    payload.Email,
		Password: payload.Password,
	}

	result := x.db.Create(&data)

	return data.UserID, result.Error
}

func (x *AccountsRepository) Read(email string) (entities.Accounts, error) {
	data := models.Accounts{}

	result := x.db.Where("email = ?", email).First(&data)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return entities.Accounts{}, common.ErrDataNotFound
	}

	account := entities.Accounts{
		UserID:   data.UserID,
		Email:    data.Email,
		Password: data.Password,
	}

	return account, result.Error
}
