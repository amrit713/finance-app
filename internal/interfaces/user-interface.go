package interfaces

import (
	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/gofiber/fiber/v2"
)

type IUserRepository interface {
	Create(user *models.User) error
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
	GetAllUsers() ([]models.User, error)
}

type IUserService interface {
	GetAllUsers() ([]models.User, error)
	GetMe(user *models.User) error
	Update(user *models.User, input *dto.UpdateUserInput) error
}
type IUserController interface {
	GetAllUsers(ctx *fiber.Ctx) error
	EditMe(ctx *fiber.Ctx) error
}
