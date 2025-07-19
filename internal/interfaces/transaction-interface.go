package interfaces

import (
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
)

type ITransitionRepository interface {
	Create(transition *models.Transaction) error
	FindByID(id string, userId uuid.UUID) (*models.Transaction, error)
	Update(transition *models.Transaction) error
	Delete(transition *models.Transaction, id string, userId uuid.UUID) error
	GetAllTranstions(userId uuid.UUID) ([]models.Transaction, error)
	GetAccountTranstions(userId uuid.UUID, accountId uuid.UUID) ([]models.Transaction, error)
}
