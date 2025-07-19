package interfaces

import (
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
)

type IBudgetRepository interface {
	Create(category *models.Budget) error
	FindByID(id string, userId *uuid.UUID) (*models.Budget, error)
	Update(category *models.Budget) error
	Delete(category *models.Budget, id string, userId uuid.UUID) error
}
