package interfaces

import (
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
)

type ICatetgoryRepository interface {
	Create(category *models.Category) error
	FindByID(id string, userId *uuid.UUID) (*models.Category, error)
	Update(category *models.Category) error
	Delete(category *models.Category, id string, userId uuid.UUID) error
	GetAllCategories(userId uuid.UUID) ([]models.Category, error) //TODO: diff account diff categories

}
