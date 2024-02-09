package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/response"
	"encoding/json"
	"net/http"
)

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.UserRole

	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		resp = response.BadRequest
		return
	}

	err = h.svc.RoleExists(inputData.RoleName)
	if err != nil {
		if err == errors.ErrDataNotFound {
			resp.Code = http.StatusBadRequest
			resp.Message = "Указанная роль не существует"
			return
		}
		resp = response.InternalServer
		return
	}

	inputData.User.Role.Name = inputData.RoleName
	err = h.svc.UpdateUser(inputData.User)
	if err != nil {
		if err == errors.ErrDataNotFound {
			resp.Code = http.StatusBadRequest
			resp.Message = "Пользователь не найден"
			return
		}
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
