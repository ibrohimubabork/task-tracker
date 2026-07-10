package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
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

func (r *TaskRepository) GetAllByID(ctx context.Context, ID uuid.UUID) (tasks []models.Tasks, err error) {
	query := `
	SELECT id, user_id, title, description, status, created_at, updated_at
	FROM tasks
	WHERE id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.DB.QueryContext(
		ctx, query, ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Tasks
		if err := rows.Scan(
			&t.ID, &t.UserID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetAllByUser(ctx context.Context, userID uuid.UUID) (tasks []models.Tasks, err error) {
	query := `
	SELECT id, user_id, title, description, status, created_at, updated_at
	FROM tasks
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.DB.QueryContext(
		ctx, query, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Tasks
		if err := rows.Scan(
			&t.ID, &t.UserID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
