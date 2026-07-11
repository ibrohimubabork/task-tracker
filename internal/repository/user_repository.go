package repository

import (
	"context"
	"database/sql"

	"github.com/ibrohimubarok/task-tracker/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepsitory(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user models.Users) error {
	query := `
		INSERT INTO users (id, email, password_hash, role)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at, updated_at
	`

	return r.DB.QueryRowContext(
		ctx, query, user.ID, user.Email, user.PasswordHash, user.Role,
	).Scan(&user.CreatedAt, &user.UpdatedAt)
}
