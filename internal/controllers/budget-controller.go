package controllers

import (
	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/gofiber/fiber/v2"
)

type BudgetController struct {
	service interfaces.IBudgetService
}

func NewBudgetController(service interfaces.IBudgetService) *BudgetController {
	return &BudgetController{service: service}
}

// GetBudget implements interfaces.IBudgetController.
func (c *BudgetController) GetBudget(ctx *fiber.Ctx) error {

	budgetId := ctx.Params("id")
	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	budget, err := c.service.GetBudget(budgetId, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Budget fetched successfully",

		"data": fiber.Map{
			"budget": budget,
		},
	})

}

// CreateBudget implements interfaces.IBudgetController.
func (c *BudgetController) CreateBudget(ctx *fiber.Ctx) error {
	var input dto.BudgetInput

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

	if input.Amount < 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Budget balance cannot be negative",
		})
	}

	budget, err := c.service.CreateBudget(&input, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{

		"success": true,
		"message": "Budget created successfully",

		"data": fiber.Map{
			"budget": budget,
		},
	})

}

// UpdateBudget implements interfaces.IBudgetController.
func (c *BudgetController) UpdateBudget(ctx *fiber.Ctx) error {
	var input dto.UpdateBudgetInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid input",
		})
	}

	budgetId := ctx.Params("id")

	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	budget, err := c.service.UpdateBudget(&input, budgetId, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"message": "Budget updated successfully",

		"data": fiber.Map{
			"budget": budget,
		},
	})

}

// DeleteBudget implements interfaces.IBudgetController.
func (c *BudgetController) DeleteBudget(ctx *fiber.Ctx) error {
	bugetId := ctx.Params("id")

	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	err := c.service.DeleteBudget(bugetId, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"message": "Budget deleted successfully",
		"data":    nil,
	})
}

var _ interfaces.IBudgetController = (*BudgetController)(nil)
