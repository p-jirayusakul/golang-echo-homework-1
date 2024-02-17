package profiles

import (
	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/repositories"
)

type profilesInteractor struct {
	profilesRepo repositories.ProfilesRepository
}

func NewProfilesInteractor(
	profilesRepo repositories.ProfilesRepository,
) *profilesInteractor {

	return &profilesInteractor{
		profilesRepo: profilesRepo,
	}
}

func (x *profilesInteractor) Create(arg entities.Profiles) (err error) {

	err = x.profilesRepo.Create(arg)
	if err != nil {
		return
	}

	return
}

func (x *profilesInteractor) Read(id string) (result entities.Profiles, err error) {

	var arg uuid.UUID
	arg.Scan(id)

	result, err = x.profilesRepo.Read(arg)
	if err != nil {
		return
	}

	return
}

func (x *profilesInteractor) Update(arg entities.UpdateProfilesDTO) (err error) {

	err = x.profilesRepo.Update(arg)
	if err != nil {
		return
	}

	return
}

func (x *profilesInteractor) Delete(arg entities.DeleteProfilesDTO) (err error) {

	err = x.profilesRepo.Delete(arg.UserID)
	if err != nil {
		return
	}

	return
}
