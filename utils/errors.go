package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidRequestBody   = errors.New("Invalid request body")
	ErrLongURLRequired      = errors.New("Long URL is required")
	ErrAPIKeyRequired       = errors.New("API key required")
	ErrInvalidAPIKey        = errors.New("Invalid API key")
	ErrURLNotFound          = errors.New("Short URL not found")
	ErrUnauthorizedAccess   = errors.New("You are not authorized to access the short URL's analytics")
	ErrNoClicksFound        = errors.New("No clicks found")
	ErrFailedToCreateURL    = errors.New("Failed to create short URL")
	ErrFailedToRetrieveData = errors.New("Failed to retrieve analytics")
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// WriteError writes an error response in JSON format.
func WriteError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}

// WriteJSON writes a JSON response.
func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
