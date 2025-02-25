package service

import (
	"context"
	"math/rand"

	"github.com/akos011221/url-shortener/models"
	"github.com/akos011221/url-shortener/storage"
)

type Shortener struct {
	db storage.Database
}

func NewShortener(db storage.Database) *Shortener {
	return &Shortener{db: db}
}

// CreateShortURL generates a short URL for the given long URL.
func (s *Shortener) CreateShortURL(ctx context.Context, longURL, tenantID string) (string, error) {
	// Generate short code (e.g., using Base62 encoding)
	shortCode := generateShortCode()

	// Store the mapping in the database
	if err := s.db.SaveURL(ctx, shortCode, longURL, tenantID); err != nil {
		return "", err
	}

	return shortCode, nil
}

// GetLongURL retrieves the long URL for the given short code.
func (s *Shortener) GetLongURL(ctx context.Context, shortCode string) (string, error) {
	return s.db.GetURL(ctx, shortCode)
}

// GetTenantByAPIKey retrieves the tenant associated with the given API key.
func (s *Shortener) GetTenantByAPIKey(ctx context.Context, apiKey string) (*models.Tenant, error) {
	return s.db.GetTenantByAPIKey(ctx, apiKey)
}

// GetURLTenantID retrieves the tenant ID associated with a short URL.
func (s *Shortener) GetURLTenantID(ctx context.Context, shortCode string) (string, error) {
	return s.db.GetURLTenantID(ctx, shortCode)
}

// // generateShortCode generates a unique short code using Base62 encoding.
func generateShortCode() string {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const codeLength = 6

	// Generate a random short code
	code := make([]byte, codeLength)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}
