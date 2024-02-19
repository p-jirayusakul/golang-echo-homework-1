package grpc_server

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	userPb "github.com/p-jirayusakul/golang-echo-homework-1/proto/_generated/users"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (x *server) CreateProfiles(ctx context.Context, req *userPb.CreateProfilesRequest) (*userPb.CreateProfilesResponse, error) {

	firstName := req.GetFirstName()
	lastName := req.GetLastName()
	phone := req.GetPhone()

	var userID uuid.UUID
	userID.Scan(req.GetUserId())

	arg := entities.Profiles{
		UserID:    userID,
		FirstName: &firstName,
		LastName:  &lastName,
		Email:     req.GetEmail(),
		Phone:     &phone,
	}

	err := x.profilesUsecase.CreateProfiles(arg)
	if err != nil {
		return nil, status.Error(codes.Code(http.StatusInternalServerError), err.Error())
	}

	return &userPb.CreateProfilesResponse{
		Status:  "success",
		Message: "success",
	}, nil
}
