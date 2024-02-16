package address

import (
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/repositories"
)

type addressInteractor struct {
	addressRepo repositories.AddressRepository
}

func NewAddressInteractor(
	addressRepo repositories.AddressRepository,
) *addressInteractor {

	return &addressInteractor{
		addressRepo: addressRepo,
	}
}

func (x *addressInteractor) Create(arg entities.Address) (err error) {

	err = x.addressRepo.Create(arg)
	if err != nil {
		return
	}

	return
}
