package mappers

import (
	userPb "github.com/p-jirayusakul/golang-echo-homework-1/proto/_generated/users"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
)

func ToProtoProfiles(x *entities.Profiles) *userPb.CreateProfilesRequest {
	return &userPb.CreateProfilesRequest{
		UserId:    x.UserID,
		FirstName: x.FirstName,
		LastName:  x.LastName,
		Email:     x.Email,
		Phone:     x.Phone,
	}
}

func ToDomainProfiles(x *userPb.CreateProfilesRequest) entities.Profiles {
	firstName := x.GetFirstName()
	lastName := x.GetLastName()
	phone := x.GetPhone()

	result := entities.Profiles{
		UserID:    x.GetUserId(),
		FirstName: &firstName,
		LastName:  &lastName,
		Email:     x.GetEmail(),
		Phone:     &phone,
	}

	return result
}
