package storage

import (
	"context"
	"testing"
)

func TestSaveAndGetURL(t *testing.T) {
	// Setup
	db, _ := NewDatabase("")

	// Test data
	shortCode := "abc123"
	longURL := "https://example.com"
	tenantID := "tenant-123"

	// Save URL
	err := db.SaveURL(context.Background(), shortCode, longURL, tenantID)
	if err != nil {
		t.Fatalf("SaveURL failed: %v", err)
	}

	// Get URL
	retrievedURL, err := db.GetURL(context.Background(), shortCode)
	if err != nil {
		t.Fatalf("GetURL failed: %v", err)
	}

	if retrievedURL != longURL {
		t.Errorf("Expected URL %s, got %s", longURL, retrievedURL)
	}
}
