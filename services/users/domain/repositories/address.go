package repositories

import "github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"

type AddressRepository interface {
	Create(payload entities.Address) error
}
