package request

type CreateProfilesReqiest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"email"`
	Phone     string `json:"phone"`
}
