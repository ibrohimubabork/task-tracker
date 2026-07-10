package app

import (
	"database/sql"

	"github.com/ibrohimubarok/task-tracker-api/internal/api/handler"
	"github.com/ibrohimubarok/task-tracker-api/internal/api/router"
	"github.com/ibrohimubarok/task-tracker-api/internal/repository"
	"github.com/ibrohimubarok/task-tracker-api/internal/service"
)

type App struct {
	Handlers *router.Handlers
}

func New(db *sql.DB) *App {
	// Repository
	taskRepo := repository.NewTaskRepository(db)

	// Service
	taskService := service.NewTaskService(taskRepo)

	// Handler
	taskHandler := handler.NewTaskHandler(taskService)
	return &App{
		Handlers: &router.Handlers{
			Task: taskHandler,
		},
	}
}
