package utils

import (
	"encoding/json"
	"net/http"
)

var (
	ErrInvalidRequestBody   = "Invalid request body"
	ErrLongURLRequired      = "Long URL is required"
	ErrAPIKeyRequired       = "API key required"
	ErrInvalidAPIKey        = "Invalid API key"
	ErrURLNotFound          = "Short URL not found"
	ErrUnauthorizedAccess   = "You are not authorized to access the short URL's analytics"
	ErrNoClicksFound        = "No clicks found"
	ErrFailedToCreateURL    = "Failed to create short URL"
	ErrFailedToRetrieveData = "Failed to retrieve analytics"
)

// WriteError writes an error response in JSON format.
func WriteError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// WriteJSON writes a JSON response.
func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
