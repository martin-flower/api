package handler

import (
	"encoding/json"
	"net/http"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		i, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		if i < 1 {
			return
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	i, err := w.Write([]byte(response))
	if err != nil {
		return
	}
	if i < 1 {
		return
	}
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
