package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/response"
	"encoding/json"
	"github.com/gorilla/context"
	"net/http"
)

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
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

	inputData.Task.Status.Name = inputData.Status.Name

	err = h.svc.UpdateTask(inputData.Task)
	if err != nil {
		if err == errors.ErrDataNotFound {
			resp.Code = 400
			resp.Message = "Задача не найдена"
			return
		} else if err == errors.ErrAccessDenied {
			resp.Code = 403
			resp.Message = "Недостаточно прав"
		} else if err == errors.ErrUserNotFound {
			resp.Code = 400
			resp.Message = "Пользователь не найден"
			return
		} else if err == errors.ErrProjectNotFound {
			resp.Code = 400
			resp.Message = "Проект не найден"
			return
		}
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
