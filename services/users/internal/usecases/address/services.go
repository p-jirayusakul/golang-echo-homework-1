package address

import (
	"github.com/google/uuid"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/repositories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/config"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/factories"
)

type addressInteractor struct {
	cfg         *config.UserConfig
	addressRepo repositories.AddressRepository
}

func NewAddressInteractor(
	config *config.UserConfig,
	dbFactory *factories.DBFactory,
) *addressInteractor {

	return &addressInteractor{
		cfg:         config,
		addressRepo: dbFactory.AddressRepo,
	}
}

func (x *addressInteractor) Create(arg entities.Address) (err error) {

	err = x.addressRepo.Create(arg)
	if err != nil {
		return
	}

	return
}

func (x *addressInteractor) Read(id string) (result entities.Address, err error) {

	var arg uuid.UUID
	arg.Scan(id)

	result, err = x.addressRepo.Read(arg)
	if err != nil {
		return
	}

	return
}

func (x *addressInteractor) Update(arg entities.Address) (err error) {

	err = x.addressRepo.Update(arg)
	if err != nil {
		return
	}

	return
}

func (x *addressInteractor) Delete(id string) (err error) {

	var arg uuid.UUID
	arg.Scan(id)

	err = x.addressRepo.Delete(arg)
	if err != nil {
		return
	}

	return
}
