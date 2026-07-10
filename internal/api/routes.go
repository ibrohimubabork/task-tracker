package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ibrohimubarok/task-tracker-api/internal/models"
	"github.com/ibrohimubarok/task-tracker-api/internal/response"
)

func Router(r chi.Router) {
	r.Get("/", Root)
}

func Root(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, models.Response[string]{
		Status: "Success",
		Data:   "Hello World!",
	})
}
