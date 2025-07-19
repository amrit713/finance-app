package repositories

import (
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// GetAllTranstions implements interfaces.ITransitionRepository.
func (r *TransactionRepository) GetAllTranstions(userId uuid.UUID) ([]models.Transaction, error) {
	var transitions []models.Transaction

	err := r.db.Where("user_id =?", userId).Find(&transitions).Error

	if err != nil {
		return nil, err
	}

	return transitions, nil
}

// GetAccountTranstions implements interfaces.ITransitionRepository.
func (r *TransactionRepository) GetAccountTranstions(userId uuid.UUID, accountId uuid.UUID) ([]models.Transaction, error) {
	var transitions []models.Transaction

	err := r.db.Where("user_id =? AND account_id:?", userId, accountId).Find(&transitions).Error

	if err != nil {
		return nil, err
	}

	return transitions, nil
}

// FindByID implements interfaces.ITransitionRepository.
func (r *TransactionRepository) FindByID(id string, userId uuid.UUID) (*models.Transaction, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var transition models.Transaction

	result := r.db.Where("id=? AND user_id=?", uid, userId).First(&transition)

	return &transition, result.Error

}

// Create implements interfaces.ITransitionRepository.
func (r *TransactionRepository) Create(transition *models.Transaction) error {
	return r.db.Create(&transition).Error
}

// Update implements interfaces.ITransitionRepository.
func (r *TransactionRepository) Update(transition *models.Transaction) error {
	return r.db.Model(transition).Where("id=? AND user_id=?", transition.ID, transition.UserID).Updates(transition).Error
}

// Delete implements interfaces.ITransitionRepository.
func (r *TransactionRepository) Delete(transition *models.Transaction, id string, userId uuid.UUID) error {
	uid, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	return r.db.Where("id=? AND user_id=?", uid, userId).Delete(&transition).Error
}

var _ interfaces.ITransitionRepository = (*TransactionRepository)(nil)
