package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ibrohimubarok/task-tracker-api/internal/models"
	"github.com/ibrohimubarok/task-tracker-api/internal/response"
	"github.com/ibrohimubarok/task-tracker-api/internal/service"
)

type TaskHandler struct {
	Service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var tasks models.Tasks
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INVALID_REQUEST",
					Message: err.Error(),
				},
			},
		})
		return
	}

	err = h.Service.Create(r.Context(), &tasks)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "SERVICE_ERROR",
					Message: err.Error(),
				},
			},
		})
		return
	}
	response.JSON(w, http.StatusCreated, models.Response[models.Tasks]{
		Status: "success",
		Data:   tasks,
	})
}
