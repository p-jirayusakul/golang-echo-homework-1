package usecases

import "github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"

type ProfilesUsecase interface {
	CreateProfiles(arg entities.Profiles) error
	Read(id string) (result entities.Profiles, err error)
	Update(arg entities.UpdateProfilesDTO) (err error)
	Delete(arg entities.DeleteProfilesDTO) (err error)
}
