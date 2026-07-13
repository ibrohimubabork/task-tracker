package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/ibrohimubarok/task-tracker/internal/models"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (r *TaskRepository) Create(ctx context.Context, userID uuid.UUID, task *models.Tasks) error {
	query := `
		INSERT INTO tasks (id, user_id, title, description, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING created_at, updated_at
	`

	return r.DB.QueryRowContext(
		ctx, query, task.ID, userID, task.Title, task.Description, task.Status,
	).Scan(&task.CreatedAt, &task.UpdatedAt)
}

func (r *TaskRepository) GetByID(ctx context.Context, ID uuid.UUID) (tasks []models.Tasks, err error) {
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

func (r *TaskRepository) Update(ctx context.Context, ID uuid.UUID, userID uuid.UUID, task *models.Tasks) error {
	query := `
		UPDATE tasks
		SET title = $1, description = $2, status = $3, updated_at = NOW()
		WHERE id = $4 AND user_id = $5
		RETURNING updated_at
	`

	err := r.DB.QueryRowContext(
		ctx, query, task.Title, task.Description, task.Status, ID, userID,
	).Scan(&task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrTaskNotFound
		}
		return err
	}
	return nil
}

func (r *TaskRepository) Delete(ctx context.Context, ID uuid.UUID, userID uuid.UUID) error {
	query := `
		DELETE FROM tasks WHERE id = $1 AND user_id = $2
	`

	result, err := r.DB.ExecContext(
		ctx, query, ID, userID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrTaskNotFound
	}

	return nil
}
