package interfaces

import (
	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IAccountRepository interface {
	Create(account *models.Account) error
	FindByID(id string, userId uuid.UUID) (*models.Account, error)
	Update(account *models.Account) error
	Delete(account *models.Account, id string, userId uuid.UUID) error
	GetAllAccounts(userId uuid.UUID) ([]models.Account, error)
}

type IAccountService interface {
	GetAllAccounts(userId uuid.UUID) ([]models.Account, error)
	GetAccount(id string, userId uuid.UUID) (*models.Account, error)
	UpdateAccount(input *dto.UpdateAccountInput, id string, userId uuid.UUID) (*models.Account, error)
	DeleteAccount(id string, userId uuid.UUID) error
	CreateAccount(input *dto.AccountInput, userId uuid.UUID) (*models.Account, error)
}

type IAccountController interface {
	DeleteAccount(ctx *fiber.Ctx) error
	UpdateAccount(ctx *fiber.Ctx) error
	CreateAccount(ctx *fiber.Ctx) error
	GetAccount(ctx *fiber.Ctx) error
	GetAllAccounts(ctx *fiber.Ctx) error
}
