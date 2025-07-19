package repositories

import (
	"fmt"

	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// GetAllAccounts implements interfaces.IAccountRepository.
func (r *AccountRepository) GetAllAccounts(userId uuid.UUID) ([]models.Account, error) {

	var accounts []models.Account

	err := r.db.Where("user_id =?", userId).Find(&accounts).Error

	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// FindByID implements interfaces.IAccountRepository.
func (r *AccountRepository) FindByID(id string, userId uuid.UUID) (*models.Account, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var account models.Account
	fmt.Print("acc", uid, userId)

	result := r.db.Where("id=? AND user_id=?", uid, userId).First(&account)

	return &account, result.Error
}

// Create implements interfaces.IAccountRepository.
func (r *AccountRepository) Create(account *models.Account) error {

	// Load related User (preload)
	// err := r.db.Preload("User").First(&account, "id = ?", account.ID).Error

	return r.db.Create(&account).Error
}

// Update implements interfaces.IAccountRepository.
func (r *AccountRepository) Update(account *models.Account) error {

	return r.db.Model(account).Where("id=? AND user_id=?", account.ID, account.UserID).Updates(account).Error
}

// Delete implements interfaces.IAccountRepository.
func (r *AccountRepository) Delete(account *models.Account, id string, userId uuid.UUID) error {

	uid, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	return r.db.Where("id=? AND user_id=?", uid, userId).Delete(&account).Error
}

var _ interfaces.IAccountRepository = (*AccountRepository)(nil)
