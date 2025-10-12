package dto

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
	Username string `json:"username" validate:"required,min=3,max=16"`
}

type RegisterResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type LoginRequest struct {
}

type LoginResponse struct {
}
