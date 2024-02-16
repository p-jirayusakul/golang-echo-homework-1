package repositories

import "github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"

type ProfilesRepository interface {
	Create(payload entities.Profiles) error
}
