package accounts

import (
	"errors"

	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/common"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/repositories"
)

type accountsInteractor struct {
	accountsRepo repositories.AccountsRepository
}

func NewAccountsInteractor(
	accountsRepo repositories.AccountsRepository,
) *accountsInteractor {

	return &accountsInteractor{
		accountsRepo: accountsRepo,
	}
}

func (x *accountsInteractor) Create(arg entities.Accounts) (id string, err error) {

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

	id = userId.String()
	return
}

func (x *accountsInteractor) Read(email string) (result entities.Accounts, err error) {

	result, err = x.accountsRepo.Read(email)
	if err != nil {
		return
	}

	return
}
