package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/response"
	"encoding/json"
	"github.com/gorilla/context"
	"net/http"
)

func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.ProjectStatus

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
			resp.Code = http.StatusBadRequest
			resp.Message = "Указанный статус не существует"
			return
		}
		resp = response.InternalServer
		return
	}

	inputData.Project.Status.Name = inputData.Status.Name

	inputData.Project.Manager.Email = inputData.ManagerEmail.Email
	err = h.svc.UpdateProject(inputData.Project)
	if err != nil {
		if err == errors.ErrDataNotFound {
			resp.Code = http.StatusBadRequest
			resp.Message = "Проект не найден"
			return
		} else if err == errors.ErrAccessDenied {
			resp.Code = 401
			resp.Message = "Доступ запрещен. Недостаточно прав для обновления проекта."
			return
		}
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
