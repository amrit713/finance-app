package controllers

import (
	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	service interfaces.ICategoryService
}

func NewCategoryController(service interfaces.ICategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

// GetAllCategories implements interfaces.ICagetoryController.
func (c *CategoryController) GetAllCategories(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}
	categories, err := c.service.GetAllCategories(user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"total_categories": len(categories),
		"success":          true,
		"message":          "Categories fetched successfully",

		"data": fiber.Map{
			"categories": categories,
		},
	})
}

// GetCategory implements interfaces.ICagetoryController.
func (c *CategoryController) GetCategory(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("id")
	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}
	category, err := c.service.GetCategory(categoryId, &user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "category fetched successfully",

		"data": fiber.Map{
			"category": category,
		},
	})
}

// CreateCategory implements interfaces.ICagetoryController.
func (c *CategoryController) CreateCategory(ctx *fiber.Ctx) error {
	var input dto.CategoryInput
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

	category, err := c.service.CreateCategory(&input, &user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{

		"success": true,
		"message": "Category created successfully",

		"data": fiber.Map{
			"category": category,
		},
	})

}

// UpdateCategory implements interfaces.ICagetoryController.
func (c *CategoryController) UpdateCategory(ctx *fiber.Ctx) error {
	var input dto.UpdateCategoryInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid input",
		})
	}

	categoryId := ctx.Params("id")

	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	category, err := c.service.UpdateCategory(&input, categoryId, &user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"message": "Category updated successfully",

		"data": fiber.Map{
			"category": category,
		},
	})

}

// DeleteCategory implements interfaces.ICagetoryController.
func (c *CategoryController) DeleteCategory(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("id")

	user, ok := ctx.Locals("user").(*models.User)

	if !ok {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user context",
		})
	}

	err := c.service.DeleteCategory(categoryId, &user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"message": "Category deleted successfully",
		"data":    nil,
	})
}

var _ interfaces.ICagetoryController = (*CategoryController)(nil)
