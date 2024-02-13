package handlers

import (
	"SUP/internal/models"
	"SUP/pkg/response"
	"encoding/json"
	"net/http"
)

func (h *Handler) DeleteProjectParticipant(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.ProjectParticipant

	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		resp = response.BadRequest
		return
	}

	err = h.svc.DeleteProjectParticipant(inputData)
	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
