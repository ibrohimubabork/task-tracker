package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/ibrohimubarok/task-tracker/internal/api/handler"
)

func TaskRoutes(r chi.Router, h *handler.TaskHandler, tokenAuth *jwtauth.JWTAuth) {

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Route("/tasks", func(r chi.Router) {
			r.Post("/", h.Create)
			r.Get("/{id}", h.GetByID)
			r.Get("/user", h.GetAllByUser)
			r.Put("/{id}", h.Update)
			r.Delete("/{id}", h.Delete)
		})
	})
}
