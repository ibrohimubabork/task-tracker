package app

import (
	"database/sql"

	"github.com/ibrohimubarok/task-tracker/internal/api/handler"
	"github.com/ibrohimubarok/task-tracker/internal/api/router"
	"github.com/ibrohimubarok/task-tracker/internal/repository"
	"github.com/ibrohimubarok/task-tracker/internal/service"
)

type App struct {
	Handlers *router.Handlers
}

func New(db *sql.DB) *App {
	// Repository
	taskRepo := repository.NewTaskRepository(db)
	userRepo := repository.NewUserRepsitory(db)

	// Service
	taskService := service.NewTaskService(taskRepo)
	userService := service.NewUserService(userRepo)

	// Handler
	taskHandler := handler.NewTaskHandler(taskService)
	userHandler := handler.NewUserHandler(userService)
	return &App{
		Handlers: &router.Handlers{
			Task: taskHandler,
			User: userHandler,
		},
	}
}
