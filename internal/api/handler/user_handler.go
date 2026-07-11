package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ibrohimubarok/task-tracker/internal/models"
	"github.com/ibrohimubarok/task-tracker/internal/response"
	"github.com/ibrohimubarok/task-tracker/internal/service"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request models.RegisterRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		response.JSON(w, http.StatusBadRequest, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INVALID_REQUEST",
					Message: "invalid JSON format",
				},
			},
		})
		return
	}

	err := h.Service.Register(r.Context(), request)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INTERNAL_SERVER_ERROR",
					Message: "failed to create user",
				},
			},
		})
		return
	}
	response.JSON(w, http.StatusCreated, models.Response[string]{
		Status: "success",
		Data:   "account has been created",
	})
}
