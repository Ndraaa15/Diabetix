package dto

type RegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Birth           string `json:"birth" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type VerificationRequest struct {
	ID   string `json:"id" validate:"required"`
	Code string `json:"code" validate:"required"`
}

type SignInResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
