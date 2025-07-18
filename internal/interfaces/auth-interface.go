package interfaces

import (
	"github.com/amirt713/finance-app/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type IAuthService interface {
	Register(input *dto.UserRequestType) (*dto.AuthResponse, error)
	Login(input *dto.LoginInput) (*dto.AuthResponse, error)
}

type IAuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
}
