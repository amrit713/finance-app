package dto

import (
	"github.com/amirt713/finance-app/internal/models"
)

type AuthResponse struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	OldPassword     string `json:"old_password,omitempty"`
	NewPassword     string `json:"new_password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}

type UserRequestType struct {
	Name     string `json:"name"`
	Email    string ` json:"email"`
	Password string `json:"password"`
	Image    string `json:"image"`
}
