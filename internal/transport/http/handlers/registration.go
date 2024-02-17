package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/response"
	"encoding/json"
	"net/http"
)

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
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
			resp.Code = 400
			resp.Message = "Указанная роль не существует"
			return
		}
		resp = response.InternalServer
		return
	}

	err = h.svc.Registration(inputData.User, inputData.RoleName)

	if err != nil {
		if err == errors.ErrAlreadyHasUser {
			resp.Code = 400
			resp.Message = "Пользователь с таким email уже существует"
			return
		}

		resp = response.InternalServer

		return
	}

	resp = response.Success
}
