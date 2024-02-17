package usecases

import "github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"

type ResetPasswordUsecase interface {
	Create(arg entities.ResetPassword) (id string, err error)
	Read(requestId string) (result entities.ResetPassword, err error)
}
