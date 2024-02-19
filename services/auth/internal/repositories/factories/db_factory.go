package factories

import (
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/repositories"
	db_repositories "github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/repositories/db"
	"gorm.io/gorm"
)

type DBFactory struct {
	AccountsRepo      repositories.AccountsRepository
	ResetPasswordRepo repositories.ResetPasswordRepository
}

func NewDBFactory(db *gorm.DB) *DBFactory {
	var (
		AccountsRepo      = db_repositories.NewAccountRepository(db)
		ResetPasswordRepo = db_repositories.NewResetPasswordRepository(db)
	)

	return &DBFactory{
		AccountsRepo:      &AccountsRepo,
		ResetPasswordRepo: &ResetPasswordRepo,
	}
}
