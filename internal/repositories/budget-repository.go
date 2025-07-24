package repositories

import (
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BudgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) *BudgetRepository {
	return &BudgetRepository{db: db}
}

// FindByID implements interfaces.IBudgetRepository.
func (r *BudgetRepository) FindByID(id string, userId uuid.UUID) (*models.Budget, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var budget models.Budget

	result := r.db.Where("id=? AND user_id=?", uid, userId).First(&budget)

	return &budget, result.Error

}

func (r *BudgetRepository) FindByUserID(userId uuid.UUID) (*models.Budget, error) {
	var budget models.Budget

	result := r.db.Where("user_id=?", userId).First(&budget)

	return &budget, result.Error
}

// Create implements interfaces.IBudgetRepository.
func (r *BudgetRepository) Create(budget *models.Budget) error {

	return r.db.Create(&budget).Error
}

// Update implements interfaces.IBudgetRepository.
func (r *BudgetRepository) Update(budget *models.Budget) error {
	return r.db.Model(budget).Where("id=? AND user_id=?", budget.ID, budget.UserID).Updates(budget).Error
}

// Delete implements interfaces.IBudgetRepository.
func (r *BudgetRepository) Delete(budget *models.Budget, id string, userId uuid.UUID) error {
	uid, err := uuid.Parse(id)

	if err != nil {
		return err
	}
	return r.db.Where("id=? AND user_id=?", uid, userId).Delete(&budget).Error
}

var _ interfaces.IBudgetRepository = (*BudgetRepository)(nil)
