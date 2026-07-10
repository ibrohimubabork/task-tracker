package repository

import (
	"context"
	"database/sql"

	"github.com/ibrohimubarok/task-tracker-api/internal/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (r *TaskRepository) Create(ctx context.Context, task *models.Tasks) error {
	query := `
	INSERT INTO tasks (id, user_id, title, description, status)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING created_at, updated_at
	`

	return r.DB.QueryRowContext(
		ctx, query, task.ID, task.UserID, task.Title, task.Description, task.Status,
	).Scan(&task.CreatedAt, &task.UpdatedAt)
}
