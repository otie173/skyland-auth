package services

import (
	"fmt"

	"github.com/otie173/skyland-auth/internal/domain/models"
	"github.com/otie173/skyland-auth/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
}

func NewAuthService(
	userRepo repositories.UserRepository,
	tokenRepo repositories.TokenRepository,
) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (s *AuthService) Register(username string, email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("cant hash password: %w", err)
	}

	userModel := models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	if err := s.userRepo.Create(userModel); err != nil {
		return fmt.Errorf("cant create user in user repository: %w", err)
	}

	return nil
}
