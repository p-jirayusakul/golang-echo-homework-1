package request

type CreateProfilesRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"email"`
	Phone     string `json:"phone"`
}
