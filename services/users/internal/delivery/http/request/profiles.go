package request

type CreateProfilesRequest struct {
	UserID    string `json:"userId" validate:"uuid4"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"email"`
	Phone     string `json:"phone"`
}

type FindProfilesByUserId struct {
	UserID string `param:"user_id" validate:"uuid4"`
}

type UpdateProfilesRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}
