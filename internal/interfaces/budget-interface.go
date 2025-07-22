package interfaces

import (
	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
)

type IBudgetRepository interface {
	Create(category *models.Budget) error
	FindByID(id string, userId uuid.UUID) (*models.Budget, error)
	Update(category *models.Budget) error
	Delete(category *models.Budget, id string, userId uuid.UUID) error
}

type IBudgetService interface {
	GetBudget(id string, userId uuid.UUID) (*models.Budget, error)
	UpdateBudget(input *dto.UpdateBudgetInput, id string, userId uuid.UUID) (*models.Budget, error)
	DeleteBudget(id string, userId uuid.UUID) error
	CreateBudget(input *dto.BudgetInput, userId uuid.UUID) (*models.Budget, error)
}
