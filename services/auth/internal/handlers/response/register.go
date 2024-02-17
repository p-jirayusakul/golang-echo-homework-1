package response

type RegisterResponse struct {
	UserID string `json:"userId"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
