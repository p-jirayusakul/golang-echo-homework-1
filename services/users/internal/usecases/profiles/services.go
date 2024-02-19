package profiles

import (
	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/repositories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/config"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/factories"
)

type profilesInteractor struct {
	cfg          *config.UserConfig
	profilesRepo repositories.ProfilesRepository
}

func NewProfilesInteractor(
	config *config.UserConfig,
	dbFactory *factories.DBFactory,
) *profilesInteractor {

	return &profilesInteractor{
		cfg:          config,
		profilesRepo: dbFactory.ProfilesRepo,
	}
}

func (x *profilesInteractor) CreateProfiles(arg entities.Profiles) (err error) {

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
