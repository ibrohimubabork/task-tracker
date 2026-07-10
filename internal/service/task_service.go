package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ibrohimubarok/task-tracker-api/internal/models"
	"github.com/ibrohimubarok/task-tracker-api/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		Repo: repo,
	}
}

func (s *TaskService) Create(ctx context.Context, task *models.Tasks) error {
	task.ID = uuid.New()

	return s.Repo.Create(ctx, task)
}

func (s *TaskService) GetAllByID(ctx context.Context, ID uuid.UUID) ([]models.Tasks, error) {
	return s.Repo.GetAllByID(ctx, ID)
}

func (s *TaskService) GetAllByUser(ctx context.Context, userID uuid.UUID) ([]models.Tasks, error) {
	return s.Repo.GetAllByUser(ctx, userID)
}
