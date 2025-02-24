package service

import (
	"context"

	//"github.com/akos011221/url-shortener/models"
	"github.com/akos011221/url-shortener/storage"
)

type Shortener struct {
	db storage.Database
}

func NewShortener(db storage.Database) *Shortener {
	return &Shortener{db: db}
}

// CreateShortURL generates a short URL for the given long URL.
func (s *Shortener) CreateShortURL(ctx context.Context, longURL string) (string, error) {
	// Generate short code (e.g., using Base62 encoding)
	shortCode := generateShortCode()

	// Store the mapping in the database
	if err := s.db.SaveURL(ctx, shortCode, longURL); err != nil {
		return "", err
	}

	return shortCode, nil
}

// GetLongURL retrieves the long URL for the given short code.
func (s *Shortener) GetLongURL(ctx context.Context, shortCode string) (string, error) {
	return s.db.GetURL(ctx, shortCode)
}

// generateShortCode generates a unique short code.
func generateShortCode() string {
	// TODO: Implement Base62 encoding or use a UUID library
	return "abc123"
}
