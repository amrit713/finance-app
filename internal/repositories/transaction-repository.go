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
func (r *TransactionRepository) GetAllTransactions(userId uuid.UUID) ([]models.Transaction, error) {
	var transitions []models.Transaction

	err := r.db.Where("user_id =?", userId).Find(&transitions).Error

	if err != nil {
		return nil, err
	}

	return transitions, nil
}

// GetAccountTranstions implements interfaces.ITransitionRepository.
func (r *TransactionRepository) GetAccountTransactions(userId uuid.UUID, accountId uuid.UUID) ([]models.Transaction, error) {
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

	var transaction models.Transaction

	result := r.db.Where("id=? AND user_id=?", uid, userId).First(&transaction)

	return &transaction, result.Error

}

// Create implements interfaces.ITransitionRepository.
func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(&transaction).Error
}

// Update implements interfaces.ITransitionRepository.
func (r *TransactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Model(transaction).Where("id=? AND user_id=?", transaction.ID, transaction.UserID).Updates(transaction).Error
}

// Delete implements interfaces.ITransitionRepository.
func (r *TransactionRepository) Delete(transaction *models.Transaction, id string, userId uuid.UUID) error {
	uid, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	return r.db.Where("id=? AND user_id=?", uid, userId).Delete(&transaction).Error
}

var _ interfaces.ITransactionRepository = (*TransactionRepository)(nil)
