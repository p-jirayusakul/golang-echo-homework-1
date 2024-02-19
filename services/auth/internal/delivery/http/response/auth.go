package response

type RegisterResponse struct {
	ID string `json:"id"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
