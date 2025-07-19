package controllers

import (
	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService interfaces.IUserService
}

func NewUserController(service interfaces.IUserService) *UserController {
	return &UserController{userService: service}
}

// GetAllUsers implements interfaces.IUserController.
func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := c.userService.GetAllUsers()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false,
			"error": "Failed to fetch users",
		})
	}

	return ctx.JSON(fiber.Map{
		"success":     true,
		"total_users": len(users),
		"message":     "User fetched successfully",

		"data": fiber.Map{
			"users": users,
		},
	})
}

// EditMe implements interfaces.IUserController.
func (c *UserController) EditMe(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("user").(*models.User)
	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	var input dto.UpdateUserInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,
			"error": "Invalid input",
		})
	}
	err := c.userService.Update(user, &input)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",

		"data": fiber.Map{
			"user": user,
		},
	})
}

var _ interfaces.IUserController = (*UserController)(nil)
