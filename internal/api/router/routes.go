package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ibrohimubarok/task-tracker/internal/api/handler"
	"github.com/ibrohimubarok/task-tracker/internal/models"
	"github.com/ibrohimubarok/task-tracker/internal/response"
)

type Handlers struct {
	Task *handler.TaskHandler
	User *handler.UserHandler
}

func Route(r chi.Router, h *Handlers) {
	r.Get("/", Root)
	TaskRoutes(r, h.Task)
	UserRoutes(r, h.User)
}

func Root(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, models.Response[string]{
		Status: "Success",
		Data:   "Hello World!",
	})
}
