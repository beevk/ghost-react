package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type APIError struct {
	Error string
}

type APIFunc func(http.ResponseWriter, *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func MakeHttpHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			jsonErr := WriteJSON(w, http.StatusBadRequest, APIError{
				Error: err.Error(),
			})
			if jsonErr != nil {
				fmt.Println("Failed to Encode to JSON")
			}
		}
	}
}

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
