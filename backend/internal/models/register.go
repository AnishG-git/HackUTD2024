package models

type RegisterRequest struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type RegisterResponse struct {
	Result string `json:"result"`
}
