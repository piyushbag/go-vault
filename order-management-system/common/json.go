package common

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	(w).Header().Set("Content-Type", "application/json")
	(w).WriteHeader(status)
	// encode data as JSON
	json.NewEncoder(w).Encode(data)
	// write data to response
}

func ReadJSON(r *http.Request, data interface{}) error {
	// decode JSON from request body
	return json.NewDecoder(r.Body).Decode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	// write error response
	WriteJSON(w, status, map[string]string{"error": message})
}
