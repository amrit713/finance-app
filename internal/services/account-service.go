package services

import (
	"errors"

	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
)

type AccountService struct {
	repo interfaces.IAccountRepository
}

func NewAccountService(repo interfaces.IAccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

// GetAllAccounts implements interfaces.IAccountService.
func (s *AccountService) GetAllAccounts(userId uuid.UUID) ([]models.Account, error) {

	return s.repo.GetAllAccounts(userId)
}

// GetAccount implements interfaces.IAccountService.
func (s *AccountService) GetAccount(id string, userId uuid.UUID) (*models.Account, error) {
	account, err := s.repo.FindByID(id, userId)

	if err != nil {
		return nil, errors.New("account not found")
	}

	return account, nil
}

// CreateAccount implements interfaces.IAccountService.
func (s *AccountService) CreateAccount(input *dto.AccountInput, userId uuid.UUID) (*models.Account, error) {

	account := &models.Account{
		Name:    input.Name,
		Type:    input.Type,
		Balance: input.Balance,
		UserID:  userId,
	}

	if input.IsDefault != nil {
		account.IsDefault = *input.IsDefault
	}

	err := s.repo.Create(account)

	if err != nil {
		return nil, errors.New("unable to create account")
	}

	return account, nil

}

// UpdateAccount implements interfaces.IAccountService.
func (s *AccountService) UpdateAccount(input *dto.UpdateAccountInput, id string, userId uuid.UUID) (*models.Account, error) {
	account, err := s.repo.FindByID(id, userId)

	if err != nil {
		return nil, errors.New("invalid account for this user")
	}
	if input.Name != nil {
		account.Name = *input.Name
	}
	if input.Balance != nil {
		account.Balance = *input.Balance
	}

	if input.Type != nil {
		account.Type = *input.Type
	}

	if input.IsDefault != nil {
		account.IsDefault = *input.IsDefault
	}

	err = s.repo.Update(account)

	if err != nil {
		return nil, errors.New("unable to update account")
	}

	return account, nil

}

// DeleteAccounts implements interfaces.IAccountService.
func (s *AccountService) DeleteAccount(id string, userId uuid.UUID) error {
	account, err := s.repo.FindByID(id, userId)

	if err != nil {
		return errors.New("invalid account for this user")
	}

	return s.repo.Delete(account, id, userId)
}

var _ interfaces.IAccountService = (*AccountService)(nil)
