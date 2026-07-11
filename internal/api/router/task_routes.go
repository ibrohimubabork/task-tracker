package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/ibrohimubarok/task-tracker/internal/api/handler"
)

func TaskRoutes(r chi.Router, h *handler.TaskHandler) {
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Get("/{id}", h.GetByID)
		r.Get("/user/{user_id}", h.GetAllByUser)
		r.Put("/{id}/{user_id}", h.Update)
		r.Delete("/{id}/{user_id}", h.Delete)
	})
}
