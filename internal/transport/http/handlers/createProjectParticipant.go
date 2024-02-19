package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/response"
	"encoding/json"
	"github.com/gorilla/context"
	"net/http"
)

func (h *Handler) CreateProjectParticipant(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.ProjectParticipant

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

	err = h.svc.CreateProjectParticipant(inputData)

	if err != nil {
		if err == errors.ErrDataNotFound {
			resp = response.BadRequest
			return
		} else if err == errors.ErrUserNotFound {
			resp.Code = 400
			resp.Message = "Участник не найден"
			return
		} else if err == errors.ErrProjectNotFound {
			resp.Code = 400
			resp.Message = "Проект не найден"
			return
		} else if err == errors.ErrAlreadyHasUser {
			resp.Code = 409
			resp.Message = "Участник уже есть"
			return
		}
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
