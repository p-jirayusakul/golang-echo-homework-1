package usecases

import "github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"

type AccountsUsecase interface {
	Create(arg entities.Accounts) error
}
