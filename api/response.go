package api

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

func JSONResponse(w http.ResponseWriter, status int, message string, detail interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(
		ApiResponse{
			Message: message,
			Detail:  detail,
		},
	)
}
