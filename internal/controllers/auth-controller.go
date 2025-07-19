package controllers

import (
	"time"

	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service interfaces.IAuthService
}

func NewAuthController(service interfaces.IAuthService) *AuthController {
	return &AuthController{service: service}
}

// Register implements interfaces.IAuthController.
func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var input dto.UserRequestType
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid input"})
	}

	user, err := c.service.Register(&input)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	utils.SetJwtCookie(ctx, user.Token)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User register  successfully",

		"data": fiber.Map{
			"user": dto.AuthResponse{User: user.User, Token: user.Token},
		},
	})
}

// Login implements interfaces.IAuthController.
func (c *AuthController) Login(ctx *fiber.Ctx) error {

	var input dto.LoginInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,
			"error": "Invalid input format",
		})
	}

	user, err := c.service.Login(&input)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	utils.SetJwtCookie(ctx, user.Token)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"message": "User login  successfully",

		"data": fiber.Map{
			"user": dto.AuthResponse{User: user.User, Token: user.Token},
		},
	})
}

// Logout implements interfaces.IAuthController.
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt_token",
		Value:    "",                             // Clear the value
		Expires:  time.Now().Add(-1 * time.Hour), // Expire the cookie
		MaxAge:   -1,                             // Expire immediately
		Path:     "/",                            // Match original cookie path
		HTTPOnly: true,
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User logout successfully",

		"data": nil,
	})
}

var _ interfaces.IAuthController = (*AuthController)(nil)
