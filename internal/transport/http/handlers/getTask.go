package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/response"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	var inputData models.Task
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		resp = response.BadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	taskFromDB, err := h.svc.GetTask(inputData)
	if err != nil {
		log.Println("Error:", err) // Логгирование ошибки
		if err == errors.ErrDataNotFound {
			resp = response.Response{Code: 404, Message: "Задача не найдена"}
		} else {
			resp = response.InternalServer
		}
	} else {
		resp = response.Success
		resp.Payload = map[string]interface{}{
			"id":            taskFromDB.Id,
			"title":         taskFromDB.Title,
			"description":   taskFromDB.Description,
			"status_id":     taskFromDB.Status.Id,
			"controller_id": taskFromDB.Controller.Id,
			"executor_id":   taskFromDB.Executor.Id,
			"project_id":    taskFromDB.Project.Id,
			"start_date":    taskFromDB.StartDate,
			"dead_line":     taskFromDB.DeadLine,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
