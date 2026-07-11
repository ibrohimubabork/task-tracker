package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/ibrohimubarok/task-tracker/internal/api/handler"
)

func UserRoutes(r chi.Router, h *handler.UserHandler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", h.Register)
	})
}
