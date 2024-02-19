package accounts

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/common"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/repositories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/repositories/factories"
)

type accountsInteractor struct {
	accountsRepo      repositories.AccountsRepository
	resetPasswordRepo repositories.ResetPasswordRepository
	usersGrpcRepo     repositories.UsersRepository
}

func NewAccountsInteractor(
	accountsRepo repositories.AccountsRepository,
	resetPasswordRepo repositories.ResetPasswordRepository,
	grpcFactory *factories.GrpcClientFactory,
) *accountsInteractor {

	return &accountsInteractor{
		accountsRepo:      accountsRepo,
		resetPasswordRepo: resetPasswordRepo,
		usersGrpcRepo:     grpcFactory.UsersRepo,
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

	// create profilse
	profiles := entities.Profiles{
		UserID: id,
		Email:  arg.Email,
	}

	err = x.usersGrpcRepo.CreateProfiles(context.Background(), &profiles)
	if err != nil {
		return
	}

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
