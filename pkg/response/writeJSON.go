package response

import (
	"encoding/json"
	"net/http"
)

func (r *Response) WriteJSON(w http.ResponseWriter) {
	if r.Code == 0 {
		r.Code = 200
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(r.Code)

	responseBody, err := json.Marshal(r)

	if err != nil {
		return
	}

	w.Write(responseBody)
}
