package service

import (
	"context"
	"testing"

	"github.com/akos011221/url-shortener/storage"
)

func TestCreateShortURL(t *testing.T) {
	// Setup
	db, _ := storage.NewDatabase("")
	shortener := NewShortener(db)

	// Test data
	longURL := "https://example.com"
	tenantID := "tenant-123"

	// Call method
	shortURL, err := shortener.CreateShortURL(context.Background(), longURL, tenantID)
	if err != nil {
		t.Fatalf("CreateShortURL failed: %v", err)
	}

	if shortURL == "" {
		t.Error("Expected a short URL, got an empty string")
	}
}
