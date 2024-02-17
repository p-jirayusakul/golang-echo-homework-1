package request

type RequestResetPasswordRequest struct {
	Email string `json:"email" validate:"email,required"`
}

type ResetPasswordRequest struct {
	RequestID string `json:"requestId"`
	Password  string `json:"password"`
}
