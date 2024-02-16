package usecases

import "github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"

type ProfilesUsecase interface {
	Create(arg entities.Profiles) error
}
