package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akos011221/url-shortener/storage"
)

func TestAuthMiddleware(t *testing.T) {
	// Setup
	db, _ := storage.NewDatabase("")
	middleware := AuthMiddleware(db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Test request with valid API key
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("x-api-key", "api-key-123")

	w := httptest.NewRecorder()
	middleware.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}

	// Test request with invalid API key
	req = httptest.NewRequest("GET", "/", nil)
	req.Header.Set("x-api-key", "invalid-key")

	w = httptest.NewRecorder()
	middleware.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status ode %d, got %d", http.StatusUnauthorized, w.Result().StatusCode)
	}
}
