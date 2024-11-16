package models

type LoginRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type LoginResponse struct {
	JWT string `json:"jwt"`
}