package entities

type Profiles struct {
	UserID    string  `json:"userId"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     string  `json:"email"`
	Phone     *string `json:"phone"`
}
