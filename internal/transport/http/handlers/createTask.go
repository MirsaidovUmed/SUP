package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/response"
	"encoding/json"
	"github.com/gorilla/context"
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

	_, ok := context.Get(r, "role_id").(int64)

	if !ok {
		resp.Code = 400
		resp.Message = "Не удалось получить значение role_id из контекста"
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
		if err == errors.ErrTaskAlreadyExists {
			resp.Code = 409
			resp.Message = "Задача с таким именем существует"
			return
		} else if err == errors.ErrAccessDenied {
			resp = response.Forbidden
			return
		} else if err == errors.ErrDataNotFound {
			resp = response.BadRequest
			return
		}
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
