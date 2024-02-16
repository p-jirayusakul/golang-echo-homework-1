package profiles

import (
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
