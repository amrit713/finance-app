package services

import (
	"errors"

	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/amirt713/finance-app/internal/utils"
)

type UserService struct {
	repo interfaces.IUserRepository
}

func NewUserService(repo interfaces.IUserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetMe implements interfaces.IUserService.
func (s *UserService) GetMe(user *models.User) error {
	panic("unimplemented")
}

// GetAllUsers implements interfaces.IUserService.
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

// Update implements interfaces.IUserService.
func (s *UserService) Update(user *models.User, input *dto.UpdateUserInput) error {
	// TODO: TRIM SPACE IN FIRST AND LAST IN SENTENCE
	//update name/email

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}

	//change

	if input.OldPassword != "" && input.NewPassword != "" {
		//match old password

		if !utils.ComparePasswordHash(input.OldPassword, user.Password) {
			return errors.New("old password is incorrect")
		}

		if input.NewPassword != input.ConfirmPassword {
			return errors.New("new password and confirm password do not match")
		}

		hashed, err := utils.HashPassword(input.NewPassword)

		if err != nil {
			return errors.New("failed to hashed password")
		}

		user.Password = hashed
	}

	return s.repo.Update(user)
}

var _ interfaces.IUserService = (*UserService)(nil)
