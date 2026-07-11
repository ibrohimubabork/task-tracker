package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/ibrohimubarok/task-tracker/internal/models"
	"github.com/ibrohimubarok/task-tracker/internal/repository"
	"github.com/ibrohimubarok/task-tracker/internal/security"
)

var (
	ErrEmailRequired    = errors.New("email is required")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordTooShort = errors.New("password must be at least 8 characters")
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, request models.RegisterRequest) error {
	var user models.Users
	email := strings.TrimSpace(strings.ToLower(request.Email))
	passwordHash, err := security.HashPassword(request.Password)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}
	user.ID = uuid.New()
	user.Email = email
	user.PasswordHash = passwordHash
	user.Role = "user"
	return s.Repo.Create(ctx, user)
}
