package interfaces

import (
	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ICategoryRepository interface {
	Create(category *models.Category) error
	FindByID(id string, userId *uuid.UUID) (*models.Category, error)
	Update(category *models.Category) error
	Delete(category *models.Category, id string, userId uuid.UUID) error
	GetAllCategories(userId uuid.UUID) ([]models.Category, error) //TODO: diff account diff categories

}

type ICategoryService interface {
	GetAllCategories(userId uuid.UUID) ([]models.Category, error)
	GetCategory(id string, userId *uuid.UUID) (*models.Category, error)
	CreateCategory(input *dto.CategoryInput, userId *uuid.UUID) (*models.Category, error)
	UpdateCategory(input *dto.UpdateCategoryInput, id string, userId *uuid.UUID) (*models.Category, error)
	DeleteCategory(id string, userId *uuid.UUID) error
}

type ICagetoryController interface {
	DeleteCategory(ctx *fiber.Ctx) error
	UpdateCategory(ctx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
	GetCategory(ctx *fiber.Ctx) error
	GetAllCategories(ctx *fiber.Ctx) error
}
