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
	inputData.User.Role.Name = inputData.RoleName
	err = h.svc.UpdateUser(inputData.User)
	if err != nil {
		if err == errors.ErrDataNotFound {
			resp.Code = 400
			resp.Message = "Пользователь не найден"
			return
		} else if err == errors.ErrChangeRole {
			resp.Code = 422
			resp.Message = "Нельзя изменить роль"
			return
		}
		resp = response.InternalServer
		return
	}
	resp = response.Success
}
