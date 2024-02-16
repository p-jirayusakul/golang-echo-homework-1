package usecases

import "github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"

type AddressUsecase interface {
	Create(arg entities.Address) (err error)
}
