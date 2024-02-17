package usecases

import "github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"

type AccountsUsecase interface {
	Create(arg entities.Accounts) (id string, err error)
	Read(email string) (result entities.Accounts, err error)
	UpdatePassword(arg entities.UpdatePasswordAccountDTO) (err error)
}
