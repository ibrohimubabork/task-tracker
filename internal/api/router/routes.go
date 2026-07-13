package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/ibrohimubarok/task-tracker/internal/api/handler"
	"github.com/ibrohimubarok/task-tracker/internal/models"
	"github.com/ibrohimubarok/task-tracker/internal/response"
)

type Handlers struct {
	Task *handler.TaskHandler
	User *handler.UserHandler
}

func Route(r chi.Router, h *Handlers, tokenAuth *jwtauth.JWTAuth) {
	r.Get("/", Root)
	TaskRoutes(r, h.Task, tokenAuth)
	UserRoutes(r, h.User)
}

func Root(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, models.Response[string]{
		Status: "Success",
		Data:   "Hello World!",
	})
}
