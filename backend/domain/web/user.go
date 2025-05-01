package web

import "time"

type UserCreateRequest struct {
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required,email,min=1"`
	ETHAddress     string `json:"eth_address" validate:"required"`
	Role           string `json:"role" validate:"omitempty"`
	Telp           string `json:"telp" validate:"required"`
	Password       string `json:"password" validate:"required,min=8"`
	VerifyPassword string `json:"verify_password" validate:"required,min=8"`
}

type UserCreateResponse struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=1"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserLogoutResponse struct {
	Id          string    `json:"id"`
	LoggedOutAt time.Time `json:"logged_out_at"`
}

type UserGetResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	ETHAddress string    `json:"eth_addr"`
	Role       string    `json:"role"`
	Telp       string    `json:"telp"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserUpdateRequest struct {
	Email string `json:"email" validate:"required,email,min=1"`
	Name  string `json:"name" validate:"required"`
	Telp  string `json:"telp" validate:"required;"`
}

type UserUpdatePasswordRequest struct {
	Password       string `json:"password" validate:"required,min=8"`
	NewPassword    string `json:"new_password" validate:"required,min=8"`
	VerifyPassword string `json:"verify_password" validate:"required,min=8"`
}

type GetCountUserResponse struct {
	TotalUser int64 `json:"total_user"`
}
