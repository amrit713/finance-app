package interfaces

import (
	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ITransactionRepository interface {
	Create(transition *models.Transaction) error
	FindByID(id string, userId uuid.UUID) (*models.Transaction, error)
	Update(transition *models.Transaction) error
	Delete(transition *models.Transaction, id string, userId uuid.UUID) error
	GetAllTransactions(userId uuid.UUID) ([]models.Transaction, error)
	GetAccountTransactions(userId uuid.UUID, accountId string) ([]models.Transaction, error)
}

type ITransactionService interface {
	GetAllTransctions(userId uuid.UUID) ([]models.Transaction, error)
	GetAccountTransactions(userId uuid.UUID, accountId string) ([]models.Transaction, error)
	GetTransaction(id string, userId uuid.UUID) (*models.Transaction, error)
	CreateTransaction(input *dto.TransactionInput, userId uuid.UUID) (*models.Transaction, error)
	UpdateTransaction(input *dto.UpdateTransactionInput, id string, userId uuid.UUID) (*models.Transaction, error)
	DeleteTransaction(id string, userId uuid.UUID) error
}

type ITransactionController interface {
	DeleteTransaction(ctx *fiber.Ctx) error
	UpdateTransaction(ctx *fiber.Ctx) error
	CreateTransaction(ctx *fiber.Ctx) error
	GetTransaction(ctx *fiber.Ctx) error
	GetAllTransactions(ctx *fiber.Ctx) error
	GetAccountTransactions(ctx *fiber.Ctx) error
}
