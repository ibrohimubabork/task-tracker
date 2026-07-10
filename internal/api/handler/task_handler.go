package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
					Message: "invalid json format: " + err.Error(),
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
					Code:    "INTERNAL_SERVER_ERROR",
					Message: "failed to create tasks: " + err.Error(),
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

func (h *TaskHandler) GetAllByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	IDParsed, err := uuid.Parse(ID)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INVALID_ID",
					Message: "invalid UUID format: " + err.Error(),
				},
			},
		})
		return
	}
	tasks, err := h.Service.GetAllByID(r.Context(), IDParsed)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INTERNAL_SERVER_ERROR",
					Message: "failed to get tasks: " + err.Error(),
				},
			},
		})
		return
	}

	response.JSON(w, http.StatusCreated, models.Response[[]models.Tasks]{
		Status: "success",
		Data:   tasks,
	})
}

func (h *TaskHandler) GetAllByUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	userIDParsed, err := uuid.Parse(userID)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INVALID_USER_ID",
					Message: "invalid UUID format: " + err.Error(),
				},
			},
		})
		return
	}
	tasks, err := h.Service.GetAllByUser(r.Context(), userIDParsed)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INTERNAL_SERVER_ERROR",
					Message: "failed to get tasks: " + err.Error(),
				},
			},
		})
		return
	}

	response.JSON(w, http.StatusCreated, models.Response[[]models.Tasks]{
		Status: "success",
		Data:   tasks,
	})
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	IDParsed, err := uuid.Parse(ID)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INVALID_ID",
					Message: "invalid UUID format: " + err.Error(),
				},
			},
		})
		return
	}

	userID := chi.URLParam(r, "user_id")
	userIDParsed, err := uuid.Parse(userID)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INVALID_USER_ID",
					Message: "invalid UUID format: " + err.Error(),
				},
			},
		})
		return
	}

	var tasks models.Tasks
	err = json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INVALID_REQUEST",
					Message: "invalid json format: " + err.Error(),
				},
			},
		})
		return
	}

	err = h.Service.Update(r.Context(), IDParsed, userIDParsed, &tasks)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INTERNAL_SERVER_ERROR",
					Message: "failed to update tasks: " + err.Error(),
				},
			},
		})
		return
	}
	response.JSON(w, http.StatusCreated, models.Response[string]{
		Status: "success",
		Data:   ID + " is deleted",
	})
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	IDParsed, err := uuid.Parse(ID)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INVALID_ID",
					Message: "invalid UUID format: " + err.Error(),
				},
			},
		})
		return
	}

	userID := chi.URLParam(r, "user_id")
	userIDParsed, err := uuid.Parse(userID)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INVALID_USER_ID",
					Message: "invalid UUID format: " + err.Error(),
				},
			},
		})
		return
	}

	err = h.Service.Delete(r.Context(), IDParsed, userIDParsed)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, models.Response[any]{
			Status: "error",
			Errors: []models.ErrorResponse{
				{
					Code:    "INTERNAL_SERVER_ERROR",
					Message: "failed to delete tasks: " + err.Error(),
				},
			},
		})
		return
	}
	response.JSON(w, http.StatusCreated, models.Response[string]{
		Status: "success",
		Data:   ID + " is deleted",
	})
}
