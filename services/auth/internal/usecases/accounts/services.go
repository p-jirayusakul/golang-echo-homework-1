package accounts

import (
	"errors"

	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/common"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/repositories"
)

type accountsInteractor struct {
	accountsRepo      repositories.AccountsRepository
	resetPasswordRepo repositories.ResetPasswordRepository
}

func NewAccountsInteractor(
	accountsRepo repositories.AccountsRepository,
	resetPasswordRepo repositories.ResetPasswordRepository,
) *accountsInteractor {

	return &accountsInteractor{
		accountsRepo:      accountsRepo,
		resetPasswordRepo: resetPasswordRepo,
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

func (x *accountsInteractor) UpdatePassword(arg entities.UpdatePasswordAccountDTO) (err error) {
	id, err := utils.ChiperDecrypt(arg.RequestID, configs.Config.SECRET_KEY)
	if err != nil {
		return
	}

	var requestId uuid.UUID
	requestId.Scan(id)

	resetPassword, err := x.resetPasswordRepo.Read(requestId)
	if err != nil {
		return
	}

	arg.UserID = resetPassword.UserID

	// hash password before update
	hashedPassword, err := utils.HashPassword(arg.Password)
	if err != nil {
		return
	}
	arg.Password = hashedPassword

	err = x.accountsRepo.UpdatePassword(arg)
	if err != nil {
		return
	}

	err = x.resetPasswordRepo.UpdateDone(requestId)

	return
}
