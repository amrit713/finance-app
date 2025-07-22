package services

import (
	"errors"

	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
)

type TransactionService struct {
	repo interfaces.ITransactionRepository
}

func NewTrasactionService(repo interfaces.ITransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

// GetAllTransctions implements interfaces.ITransactionService.
func (s *TransactionService) GetAllTransctions(userId uuid.UUID) ([]models.Transaction, error) {

	return s.repo.GetAllTransactions(userId)
}

// GetAccountTranstions implements interfaces.ITransactionService.
func (s *TransactionService) GetAccountTransactions(userId uuid.UUID, accountId uuid.UUID) ([]models.Transaction, error) {
	return s.repo.GetAccountTransactions(userId, accountId)
}

// GetTransaction implements interfaces.ITransactionService.
func (s *TransactionService) GetTransaction(id string, userId uuid.UUID) (*models.Transaction, error) {
	transaction, err := s.repo.FindByID(id, userId)

	if err != nil {
		return nil, errors.New("no transaction found")
	}

	return transaction, nil
}

// CreateTransaction implements interfaces.ITransactionService.
func (s *TransactionService) CreateTransaction(input *dto.TransactionInput, userId uuid.UUID) (*models.Transaction, error) {
	transaction := &models.Transaction{
		CategoryID: input.CategoryID,
		AccountID:  input.AccountID,
		UserID:     userId,

		Amount:      input.Amount,
		Date:        input.Date,
		Name:        input.Name,
		Description: input.Description,
		Type:        input.Type,
	}

	if input.ReceiptURL != nil {
		transaction.ReceiptURL = input.ReceiptURL
	}

	if input.IsRecurring != nil {
		transaction.IsRecurring = *input.IsRecurring
	}

	if input.RecurringInterval != nil {
		transaction.RecurringInterval = input.RecurringInterval
	}
	if input.NextRecurringDate != nil {
		transaction.NextRecurringDate = input.NextRecurringDate
	}

	if input.LastProcessed != nil {
		transaction.LastProcessed = input.LastProcessed
	}

	if input.Status != nil {
		transaction.Status = *input.Status
	}

	err := s.repo.Create(transaction)

	if err != nil {
		return nil, errors.New("unable to create transaction")
	}

	return transaction, nil
}

// UpdateTransaction implements interfaces.ITransitionService.
func (s *TransactionService) UpdateTransaction(input *dto.UpdateTransactionInput, id string, userId uuid.UUID) (*models.Transaction, error) {
	transaction, err := s.repo.FindByID(id, userId)

	if err != nil {
		return nil, errors.New("unable to find transaction")
	}

	if input.CategoryID != nil {
		transaction.CategoryID = *input.CategoryID
	}
	if input.Name != nil {
		transaction.Name = *input.Name
	}
	if input.Description != nil {
		transaction.Description = *input.Description
	}
	if input.Amount != nil {
		transaction.Amount = *input.Amount
	}

	if input.Date != nil {
		transaction.Date = *input.Date
	}

	if input.Type != nil {
		transaction.Type = *input.Type
	}

	if input.ReceiptURL != nil {
		transaction.ReceiptURL = input.ReceiptURL
	}

	if input.IsRecurring != nil {
		transaction.IsRecurring = *input.IsRecurring
	}

	if input.RecurringInterval != nil {
		transaction.RecurringInterval = input.RecurringInterval
	}
	if input.NextRecurringDate != nil {
		transaction.NextRecurringDate = input.NextRecurringDate
	}

	if input.LastProcessed != nil {
		transaction.LastProcessed = input.LastProcessed
	}

	if input.Status != nil {
		transaction.Status = *input.Status
	}

	err = s.repo.Update(transaction)

	if err != nil {
		return nil, errors.New("unable to update transaction")
	}
	return transaction, nil

}

// DeleteTransaction implements interfaces.ITransitionService.
func (s *TransactionService) DeleteTransaction(id string, userId uuid.UUID) error {
	transaction, err := s.repo.FindByID(id, userId)

	if err != nil {
		return errors.New("unable to find transaction")
	}

	return s.repo.Delete(transaction, id, userId)

}

var _ interfaces.ITransactionService = (*TransactionService)(nil)
