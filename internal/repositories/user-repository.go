package repositories

import (
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

var _ interfaces.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create implements interfaces.IUserRepostory
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByEmail implements interfaces.IUserRepository
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)

	return &user, result.Error
}

// Delete implements interfaces.IUserRepository.
func (r *UserRepository) Delete(id string) error {
	panic("unimplemented")
}

// FindByID implements interfaces.IUserRepository.
func (r *UserRepository) FindByID(id string) (*models.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var user models.User

	result := r.db.Where("id =?", uid).First(&user)

	return &user, result.Error
}

// GetAllUsers implements interfaces.IUserRepository.
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Update implements interfaces.IUserRepository.
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Model(user).Where("id = ?", user.ID).Updates(user).Error
}
