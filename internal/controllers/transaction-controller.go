package controllers

import (
	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	service interfaces.ITransactionService
}

func NewTransactionController(service interfaces.ITransactionService) *TransactionController {
	return &TransactionController{service: service}
}

// GetAllTransactions implements interfaces.ITransactionController.
func (c *TransactionController) GetAllTransactions(ctx *fiber.Ctx) error {

	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	transactions, err := c.service.GetAllTransctions(user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"total_transactions": len(transactions),
		"success":            true,
		"message":            "Transactions fetched successfully",

		"data": fiber.Map{
			"transactions": transactions,
		},
	})

}

// GetAccountTransactions implements interfaces.ITransactionController.
func (c *TransactionController) GetAccountTransactions(ctx *fiber.Ctx) error {
	accountId := ctx.Params("accountId")

	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	transactions, err := c.service.GetAccountTransactions(user.ID, accountId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"total_transactions": len(transactions),
		"success":            true,
		"message":            "Transactions fetched successfully",

		"data": fiber.Map{
			"transactions": transactions,
		},
	})

}

// GetTransaction implements interfaces.ITransactionController.
func (c *TransactionController) GetTransaction(ctx *fiber.Ctx) error {
	transactionId := ctx.Params("id")
	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}
	transaction, err := c.service.GetTransaction(transactionId, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Transaction fetched successfully",

		"data": fiber.Map{
			"transaction": transaction,
		},
	})

}

// CreateTransaction implements interfaces.ITransactionController.
func (c *TransactionController) CreateTransaction(ctx *fiber.Ctx) error {
	var input dto.TransactionInput
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
			"error":   err,
		})
	}

	transaction, err := c.service.CreateTransaction(&input, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{

		"success": true,
		"message": "Transaction created successfully",

		"data": fiber.Map{
			"transaction": transaction,
		},
	})

}

// UpdateTransaction implements interfaces.ITransactionController.
func (c *TransactionController) UpdateTransaction(ctx *fiber.Ctx) error {
	var input dto.UpdateTransactionInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid input",
		})
	}

	transactionId := ctx.Params("id")

	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	transaction, err := c.service.UpdateTransaction(&input, transactionId, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"message": "Transaction updated successfully",

		"data": fiber.Map{
			"transaction": transaction,
		},
	})

}

// DeleteTransaction implements interfaces.ITransactionController.
func (c *TransactionController) DeleteTransaction(ctx *fiber.Ctx) error {

	transactionId := ctx.Params("id")

	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	err := c.service.DeleteTransaction(transactionId, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"message": "Transaction deleted successfully",
		"data":    nil,
	})

}

var _ interfaces.ITransactionController = (*TransactionController)(nil)
