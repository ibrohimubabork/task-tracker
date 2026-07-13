package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/ibrohimubarok/task-tracker/internal/models"
	"github.com/ibrohimubarok/task-tracker/internal/repository"
	"github.com/ibrohimubarok/task-tracker/internal/security"
)

var (
	ErrEmailRequired      = errors.New("email is required")
	ErrPasswordRequired   = errors.New("password is required")
	ErrPasswordTooShort   = errors.New("password must be at least 8 characters")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInvalidLoginInput  = errors.New("email and password are required")
)

const accessTokenDuration = 2 * time.Hour

type UserService struct {
	Repo      *repository.UserRepository
	TokenAuth *jwtauth.JWTAuth
}

func NewUserService(repo *repository.UserRepository, tokenAuth *jwtauth.JWTAuth) *UserService {
	return &UserService{
		Repo:      repo,
		TokenAuth: tokenAuth,
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

func (s *UserService) Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error) {
	email := strings.TrimSpace(strings.ToLower(request.Email))
	password := request.Password
	user, err := s.Repo.GetByEmail(ctx, email)
	if err != nil {
		return models.LoginResponse{}, err
	}
	matched, err := security.VerifyPassword(password, user.PasswordHash)
	if err != nil {
		return models.LoginResponse{}, err
	}
	if !matched {
		return models.LoginResponse{}, ErrInvalidCredentials
	}

	claims := map[string]interface{}{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"role":    user.Role,
	}
	jwtauth.SetIssuedNow(claims)
	jwtauth.SetExpiryIn(claims, accessTokenDuration)
	_, tokenString, err := s.TokenAuth.Encode(claims)
	if err != nil {
		return models.LoginResponse{}, fmt.Errorf(
			"encode access token: %w",
			err,
		)
	}
	return models.LoginResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   int64(accessTokenDuration.Seconds()),
	}, err
}
