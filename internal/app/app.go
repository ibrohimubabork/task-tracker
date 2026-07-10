package app

import (
	"database/sql"

	"github.com/ibrohimubarok/task-tracker-api/internal/api/handler"
	"github.com/ibrohimubarok/task-tracker-api/internal/api/routes"
	"github.com/ibrohimubarok/task-tracker-api/internal/repository"
	"github.com/ibrohimubarok/task-tracker-api/internal/service"
)

type App struct {
	Handlers *routes.Handlers
}

func New(db *sql.DB) *App {
	// Repository
	taskRepo := repository.NewTaskRepository(db)

	// Service
	taskService := service.NewTaskService(taskRepo)

	// Handler
	taskHandler := handler.NewTaskHandler(taskService)
	return &App{
		Handlers: &routes.Handlers{
			Task: taskHandler,
		},
	}
}
