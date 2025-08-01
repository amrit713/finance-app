package repositories

import (
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// GetAllCategories implements interfaces.ICatetgoryRepository.
func (r *CategoryRepository) GetAllCategories(userId uuid.UUID) ([]models.Category, error) {
	var categories []models.Category

	err := r.db.Where("user_id=? OR user_id=NULL", userId).Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// FindByID implements interfaces.ICatetgoryRepository.
func (r *CategoryRepository) FindByID(id string, userId *uuid.UUID) (*models.Category, error) {
	//TODO: USER must fetch category that is created by them
	var category models.Category

	uid, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	query := r.db.Where("id = ?", uid)

	if userId != nil {
		query = r.db.Where("user_id =?", userId)
	}

	err = query.First(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

// Create implements interfaces.ICatetgoryRepository.
func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(&category).Error
}

// Update implements interfaces.ICatetgoryRepository.
func (r *CategoryRepository) Update(category *models.Category) error {
	return r.db.Model(category).Where("id=? AND user_id=?", category.ID, category.UserID).Updates(category).Error
}

// Delete implements interfaces.ICatetgoryRepository.
func (r *CategoryRepository) Delete(category *models.Category, id string, userId uuid.UUID) error {
	uid, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	return r.db.Where("id=? AND user_id=?", uid, userId).Delete(&category).Error
}

var _ interfaces.ICategoryRepository = (*CategoryRepository)(nil)
