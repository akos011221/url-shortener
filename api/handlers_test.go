package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/akos011221/url-shortener/models"
	"github.com/akos011221/url-shortener/service"
	"github.com/akos011221/url-shortener/storage"
)

func TestCreateShortURL(t *testing.T) {
	// Setup
	db, _ := storage.NewDatabase("")
	shortener := service.NewShortener(db)
	analytics := service.NewAnalytics(db)
	handlers := Handlers{
		Shortener: shortener,
		Analytics: analytics,
	}

	// Test request
	reqBody := `{"longUrl": "https://example.com"}`
	req := httptest.NewRequest("POST", "/shorten", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "api-key-123")

	// Add tenant ID to the request context
	ctx := context.WithValue(req.Context(), "tenantID", "tenant-123")
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	// Call handler
	handlers.CreateShortURL(w, req)

	// Check response
	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	var response models.CreateShortURLResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.ShortURL == "" {
		t.Error("Expected a short URL, got an empty string")
	}
}
