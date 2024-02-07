package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/errors"
	"SUP/pkg/response"
	"encoding/json"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.User

	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		resp = response.BadRequest
		return
	}

	token, err := h.svc.Login(inputData)

	if err != nil {
		if err == errors.ErrDataNotFound {
			resp = response.NotFound
			return
		} else if err == errors.ErrWrongPassword {
			resp.Code = http.StatusForbidden
			resp.Message = "Wrong Password"
			return
		} else {
			resp = response.InternalServer
			return
		}
	}

	resp = response.Success
	resp.Payload = token
}
