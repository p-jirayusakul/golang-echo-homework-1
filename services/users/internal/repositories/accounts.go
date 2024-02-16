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
	arg := models.Accounts{
		Email:    payload.Email,
		Password: payload.Password,
	}

	result := x.db.Create(&arg)

	return arg.UserID, result.Error
}

func (x *AccountsRepository) Find(email string) (entities.Accounts, error) {
	arg := models.Accounts{
		Email: email,
	}

	result := x.db.First(&arg)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return entities.Accounts{}, common.ErrDataNotFound
	}

	account := entities.Accounts{
		UserID:   arg.UserID,
		Email:    arg.Email,
		Password: arg.Password,
	}

	return account, result.Error
}
