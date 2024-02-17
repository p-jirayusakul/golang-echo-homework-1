package accounts

import (
	"errors"

	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/common"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/repositories"
)

type accountsInteractor struct {
	accountsRepo repositories.AccountsRepository
	profilesRepo repositories.ProfilesRepository
}

func NewAccountsInteractor(
	accountsRepo repositories.AccountsRepository,
	profilesRepo repositories.ProfilesRepository,
) *accountsInteractor {

	return &accountsInteractor{
		accountsRepo: accountsRepo,
		profilesRepo: profilesRepo,
	}
}

func (x *accountsInteractor) Create(arg entities.Accounts) (err error) {

	_, err = x.accountsRepo.Read(arg.Email)
	if err != nil {
		if !errors.Is(err, common.ErrDataNotFound) {
			return
		}
	} else {
		err = errors.New("this email already used")
		return
	}

	// hash password before insert
	hashedPassword, err := utils.HashPassword(arg.Password)
	if err != nil {
		return
	}
	arg.Password = hashedPassword

	userId, err := x.accountsRepo.Create(arg)
	if err != nil {
		return
	}

	var profiles entities.Profiles
	profiles.Email = arg.Email
	profiles.UserID = userId

	err = x.profilesRepo.Create(profiles)
	if err != nil {
		return
	}

	return
}
