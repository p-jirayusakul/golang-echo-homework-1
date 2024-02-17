package reset_password

import (
	"time"

	"github.com/google/uuid"

	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/repositories"
)

type resetPasswordInteractor struct {
	resetPasswordRepo repositories.ResetPasswordRepository
	accountsRepo      repositories.AccountsRepository
}

func NewResetPasswordInteractor(
	resetPasswordRepo repositories.ResetPasswordRepository,
	accountsRepo repositories.AccountsRepository,
) *resetPasswordInteractor {

	return &resetPasswordInteractor{
		resetPasswordRepo: resetPasswordRepo,
		accountsRepo:      accountsRepo,
	}
}

func (x *resetPasswordInteractor) Create(arg entities.ResetPassword) (id string, err error) {

	user, err := x.accountsRepo.Read(arg.Email)
	if err != nil {
		return
	}

	arg.UserID = user.UserID
	arg.ExpiresAt = time.Now().Add(time.Hour * 24)
	requestId, err := x.resetPasswordRepo.Create(arg)
	if err != nil {
		return
	}

	id, err = utils.ChiperEncrypt(requestId.String(), configs.Config.SECRET_KEY)
	if err != nil {
		return
	}

	return
}

func (x *resetPasswordInteractor) Read(requestId string) (result entities.ResetPassword, err error) {

	id, err := utils.ChiperDecrypt(requestId, configs.Config.SECRET_KEY)
	if err != nil {
		return
	}

	var arg uuid.UUID
	arg.Scan(id)

	result, err = x.resetPasswordRepo.Read(arg)
	if err != nil {
		return
	}

	return
}
