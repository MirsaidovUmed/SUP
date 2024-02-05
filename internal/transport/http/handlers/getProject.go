package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/response"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) GetProject(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	var inputData models.Project
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		fmt.Println(1)
		resp = response.BadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	projectFromDB, err := h.svc.GetProject(inputData)
	if err != nil {
		fmt.Println(2)
		log.Println("Error:", err) // Логгирование ошибки
		if err == errors.ErrDataNotFound {
			resp = response.Response{Code: 404, Message: "Проект не найден"}
		} else {
			resp = response.InternalServer
		}
	} else {
		resp = response.Success
		resp.Payload = map[string]interface{}{
			"id":          projectFromDB.Id,
			"name":        projectFromDB.Name,
			"description": projectFromDB.Description,
			"status_id":   projectFromDB.Status.Id,
			"manager_id":  projectFromDB.Manager.Id,
			"start_date":  projectFromDB.StartDate,
			"dead_line":   projectFromDB.DeadLine,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
