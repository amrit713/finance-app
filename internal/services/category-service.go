package services

import (
	"errors"

	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
)

type CategoryService struct {
	repo interfaces.ICatetgoryRepository
}

func NewCategoryService(repo interfaces.ICatetgoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// GetAllCategories implements interfaces.ICategoryService.
func (s *CategoryService) GetAllCategories(userId uuid.UUID) ([]models.Category, error) {
	return s.repo.GetAllCategories(userId)
}

// GetCategory implements interfaces.ICategoryService.
func (s *CategoryService) GetCategory(id string, userId *uuid.UUID) (*models.Category, error) {
	category, err := s.repo.FindByID(id, userId)

	if err != nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}

// CreateCategory implements interfaces.ICategoryService.
func (s *CategoryService) CreateCategory(input *dto.CategoryInput, userId *uuid.UUID) (*models.Category, error) {
	category := &models.Category{
		Name: input.Name,
		Icon: input.Icon,
	}
	if userId != nil {
		category.UserID = userId
	}

	err := s.repo.Create(category)

	if err != nil {
		return nil, errors.New("unable to create category")
	}
	return category, nil

}

// UpdateCategory implements interfaces.ICategoryService.
func (s *CategoryService) UpdateCategory(input *dto.UpdateCategoryInput, id string, userId *uuid.UUID) (*models.Category, error) {
	category, err := s.repo.FindByID(id, userId)

	if err != nil {
		return nil, errors.New("invalid category for this user")
	}

	if input.Name != nil {
		category.Name = *input.Name
	}

	if input.Icon != nil {
		category.Icon = *input.Icon
	}
	err = s.repo.Update(category)

	if err != nil {
		return nil, errors.New("unable to update category")
	}

	return category, nil

}

// DeleteCategory implements interfaces.ICategoryService.
func (s *CategoryService) DeleteCategory(id string, userId *uuid.UUID) error {
	category, err := s.repo.FindByID(id, userId)

	if err != nil {
		return errors.New("invalid category for this user")
	}

	return s.repo.Delete(category, id, *userId)
}

var _ interfaces.ICategoryService = (*CategoryService)(nil)
