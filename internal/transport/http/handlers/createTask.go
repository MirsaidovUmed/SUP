package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/response"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.TaskStatus

	err := json.NewDecoder(r.Body).Decode(&inputData)

	if err != nil {
		resp = response.BadRequest
		return
	}

	err = h.svc.StatusExists(inputData.Status)
	if err != nil {
		if err == errors.ErrDataNotFound {
			resp.Code = 400
			resp.Message = "Указанный статус не существует"
			return
		}

		resp = response.InternalServer
		return
	}

	err = h.svc.CreateTask(inputData.Task, inputData.Status)

	if err != nil {
		if err == errors.ErrAlreadyHasUser {
			resp.Code = 409
			resp.Message = err.Error()
			return
		}

		resp = response.InternalServer
		return
	}

	resp = response.Success
}
