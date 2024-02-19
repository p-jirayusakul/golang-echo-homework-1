package factories

import (
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/repositories"
	db_repositories "github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/db"
	"gorm.io/gorm"
)

type DBFactory struct {
	ProfilesRepo repositories.ProfilesRepository
	AddressRepo  repositories.AddressRepository
}

func NewDBFactory(db *gorm.DB) *DBFactory {
	var (
		ProfilesRepo = db_repositories.NewProfileRepository(db)
		AddressRepo  = db_repositories.NewAddressRepository(db)
	)

	return &DBFactory{
		ProfilesRepo: &ProfilesRepo,
		AddressRepo:  &AddressRepo,
	}
}
