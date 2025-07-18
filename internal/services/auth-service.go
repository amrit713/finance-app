package services

import (
	"errors"
	"fmt"

	"github.com/amirt713/finance-app/internal/dto"
	"github.com/amirt713/finance-app/internal/interfaces"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/amirt713/finance-app/internal/utils"
)

type AuthService struct {
	repo interfaces.IUserRepository
}

func NewAuthService(repo interfaces.IUserRepository) *AuthService {
	return &AuthService{repo}
}

// Register implements interfaces.IAuthService.
func (s *AuthService) Register(input *dto.UserRequestType) (*dto.AuthResponse, error) {

	_, err := s.repo.FindByEmail(input.Email)

	if err == nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(input.Password)

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Image:    input.Image,
	}

	if err != nil {
		return nil, err
	}

	err = s.repo.Create(user)

	if err != nil {
		fmt.Println("error", err)
		return nil, errors.New("unable to create user")
	}

	token, err := utils.GenerateJWT(user.ID.String(), user.Email)

	if err != nil {
		return nil, errors.New("unable to sign jwt token")
	}

	return &dto.AuthResponse{User: user, Token: token}, nil
}

// Login implements interfaces.IAuthService.
func (s *AuthService) Login(input *dto.LoginInput) (*dto.AuthResponse, error) {
	user, err := s.repo.FindByEmail(input.Email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	//compare hash
	if !utils.ComparePasswordHash(input.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID.String(), user.Email)

	if err != nil {
		return nil, errors.New("unable to sign jwt token")
	}

	return &dto.AuthResponse{User: user, Token: token}, nil
}

var _ interfaces.IAuthService = (*AuthService)(nil)
