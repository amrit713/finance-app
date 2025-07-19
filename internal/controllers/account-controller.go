package controllers

import (
	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	accountService interfaces.IAccountService
}

func NewAccountController(service interfaces.IAccountService) *AccountController {
	return &AccountController{accountService: service}
}

// GetAllAccounts implements interfaces.IAccountController.
func (c *AccountController) GetAllAccounts(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}
	accounts, err := c.accountService.GetAllAccounts(user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"total_accounts": len(accounts),
		"success":        true,
		"message":        "Accounts fetched successfully",

		"data": fiber.Map{
			"accounts": accounts,
		},
	})

}

// GetAccount implements interfaces.IAccountController.
func (c *AccountController) GetAccount(ctx *fiber.Ctx) error {
	accountId := ctx.Params("id")
	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}
	account, err := c.accountService.GetAccount(accountId, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Account fetched successfully",

		"data": fiber.Map{
			"account": account,
		},
	})

}

// CreateAccount implements interfaces.IAccountController.
func (c *AccountController) CreateAccount(ctx *fiber.Ctx) error {
	var input dto.AccountInput
	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid input",
		})
	}

	// Validate required fields
	if input.Name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Account name is required",
		})
	}
	if input.Balance < 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Account balance cannot be negative",
		})
	}

	account, err := c.accountService.CreateAccount(&input, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{

		"success": true,
		"message": "Account created successfully",

		"data": fiber.Map{
			"account": account,
		},
	})

}

// UpdateAccount implements interfaces.IAccountController.
func (c *AccountController) UpdateAccount(ctx *fiber.Ctx) error {
	var input dto.UpdateAccountInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid input",
		})
	}

	accountId := ctx.Params("id")

	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	account, err := c.accountService.UpdateAccount(&input, accountId, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"message": "Account updated successfully",

		"data": fiber.Map{
			"account": account,
		},
	})

}

// DeleteAccount implements interfaces.IAccountController.
func (c *AccountController) DeleteAccount(ctx *fiber.Ctx) error {
	accountId := ctx.Params("id")

	user := ctx.Locals("user").(*models.User)

	err := c.accountService.DeleteAccount(accountId, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"message": "Account deleted successfully",
		"data":    nil,
	})

}

var _ interfaces.IAccountController = (*AccountController)(nil)
