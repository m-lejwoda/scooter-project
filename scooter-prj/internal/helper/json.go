package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Printf("Error appeared due to json encoding")
	}
}

func ReadJSON[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var value T
	err := json.NewDecoder(r.Body).Decode(&value)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Unsupported Json data format")
		return value, err
	}
	return value, nil
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"error": message})
}
