package service

import (
	"context"
	"testing"

	"github.com/akos011221/url-shortener/storage"
)

func TestLogClick(t *testing.T) {
	// Setup
	db, _ := storage.NewDatabase("")
	analytics := NewAnalytics(db)

	// Test data
	shortCode := "abc123"
	ipAddress := "127.0.0.1"
	userAgent := "test-agent"

	// Log a click
	err := analytics.LogClick(context.Background(), shortCode, ipAddress, userAgent)
	if err != nil {
		t.Fatalf("LogClick failed: %v", err)
	}

	// Verify the click was logged
	clicks, err := db.GetClicks(context.Background(), shortCode)
	if err != nil {
		t.Fatalf("GetClicks failed: %v", err)
	}

	if len(clicks) != 1 {
		t.Errorf("Expected 1 click, got %d", len(clicks))
	}

	click := clicks[0]
	if click.ShortCode != shortCode || click.IPAddress != ipAddress || click.UserAgent != userAgent {
		t.Errorf("Click data mismatch: got %+v", click)
	}
}

func TestGetAnalytics(t *testing.T) {
	// Setup
	db, _ := storage.NewDatabase("")
	analytics := NewAnalytics(db)

	// Test data
	shortCode := "abc123"
	ipAddress := "127.0.0.1"
	userAgent := "test-agent"

	// Log a click
	err := analytics.LogClick(context.Background(), shortCode, ipAddress, userAgent)
	if err != nil {
		t.Fatalf("LogClick failed: %v", err)
	}

	// Get analytics
	response, err := analytics.GetAnalytics(context.Background(), shortCode)
	if err != nil {
		t.Fatalf("GetAnalytics failed: %v", err)
	}

	if response.Clicks != 1 {
		t.Errorf("Expected 1 click, got %d", response.Clicks)
	}

	if len(response.Details) != 1 {
		t.Errorf("Expected 1 click details, got %d", len(response.Details))
	}

	click := response.Details[0]
	if click.ShortCode != shortCode || click.IPAddress != ipAddress || click.UserAgent != userAgent {
		t.Errorf("Click data mismatch: got %+v", click)
	}
}
