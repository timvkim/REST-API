package handlers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
