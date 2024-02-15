package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/response"
	"encoding/json"
	"github.com/gorilla/context"
	"net/http"
)

func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
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
			resp.Code = 400
			resp.Message = "Указанный статус не существует"
			return
		}
		resp = response.InternalServer
		return
	}
	err = h.svc.CreateProject(inputData.Project, inputData.Status, inputData.ManagerEmail.Email)

	if err != nil {
		if err == errors.ErrAlreadyHasUser {
			resp.Code = 409
			resp.Message = err.Error()
			return
		} else if err == errors.ErrProjectAlreadyExists {
			resp.Code = 409
			resp.Message = "Проект с таким именем существует"
			return
		} else if err == errors.ErrAccessDenied {
			resp.Code = 401
			resp.Message = "Доступ запрещен. Недостаточно прав для назначения в менеджеры."
			return
		}
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
