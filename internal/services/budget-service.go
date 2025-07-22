package services

import (
	"errors"

	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
)

type BudgetService struct {
	repo interfaces.IBudgetRepository
}

func NewBudgetService(repo interfaces.IBudgetRepository) *BudgetService {
	return &BudgetService{repo: repo}
}

// GetBudget implements interfaces.IBudgetService.
func (s *BudgetService) GetBudget(id string, userId uuid.UUID) (*models.Budget, error) {
	budget, err := s.repo.FindByID(id, userId)

	if err != nil {
		return nil, errors.New("budget not found")
	}

	return budget, nil
}

// CreateBudget implements interfaces.IBudgetService.
func (s *BudgetService) CreateBudget(input *dto.BudgetInput, userId uuid.UUID) (*models.Budget, error) {
	budget := &models.Budget{
		Amount: input.Amount,
		UserID: userId,
	}

	err := s.repo.Create(budget)

	if err != nil {
		return nil, errors.New("unable to create budget")
	}

	return budget, nil

}

// UpdateBudget implements interfaces.IBudgetService.
func (s *BudgetService) UpdateBudget(input *dto.UpdateBudgetInput, id string, userId uuid.UUID) (*models.Budget, error) {
	budget, err := s.repo.FindByID(id, userId)

	if err != nil {
		return nil, errors.New("budget not found")
	}

	if input.Amount != nil {
		budget.Amount = *input.Amount
	}

	if input.LastAlertSend != nil {
		budget.LastAlertSend = input.LastAlertSend
	}

	err = s.repo.Update(budget)

	if err != nil {
		return nil, errors.New("unable to update budget")
	}

	return budget, nil

}

// DeleteBudget implements interfaces.IBudgetService.
func (s *BudgetService) DeleteBudget(id string, userId uuid.UUID) error {
	budget, err := s.repo.FindByID(id, userId)

	if err != nil {
		return errors.New("invalid budget for this user")
	}

	return s.repo.Delete(budget, id, userId)
}

var _ interfaces.IBudgetService = (*BudgetService)(nil)
