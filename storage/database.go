package storage

import (
	"context"
	"errors"

	"github.com/akos011221/url-shortener/models"
)

type Database interface {
	SaveURL(ctx context.Context, shortCode, longURL string) error
	GetURL(ctx context.Context, shortCode string) (string, error)
	SaveClick(ctx context.Context, click models.Click) error
	GetClicks(ctx context.Context, shortCode string) ([]models.Click, error)
	GetTenantByAPIKey(ctx context.Context, apiKey string) (*models.Tenant, error)
	Close() error
}

type InMemoryDatabase struct {
	urls map[string]string // shortCode -> longURL
	clicks map[string][]models.Click // shortCode -> []Click
	tenants map[string]models.Tenant // apiKey -> Tenant
}

func NewDatabase(databaseURL string) (Database, error) {
	// Initialize in-memory data
	tenants := map[string]models.Tenant{
		"api-key-123": {ID: "1", Name: "Tenant A", APIKey: "api-key-123"},
	}

	// TODO: Connect to a real database (e.g., PostgreSQL, Redis)
	return &InMemoryDatabase{
		urls: make(map[string]string),
		clicks: make(map[string][]models.Click),
		tenants: tenants,
	}, nil
}

func (db *InMemoryDatabase) SaveURL(ctx context.Context, shortCode, longURL string) error {
	db.urls[shortCode] = longURL
	return nil
}

func (db *InMemoryDatabase) GetURL(ctx context.Context, shortCode string) (string, error) {
	longURL, ok := db.urls[shortCode]
	if !ok {
		return "", errors.New("URL not found")
	}
	return longURL, nil
}

func (db *InMemoryDatabase) SaveClick(ctx context.Context, click models.Click) error {
	db.clicks[click.ShortCode] = append(db.clicks[click.ShortCode], click)
	return nil
}

func (db *InMemoryDatabase) GetClicks(ctx context.Context, shortCode string) ([]models.Click, error) {
	clicks, ok := db.clicks[shortCode]
	if !ok {
		return nil, errors.New("No clicks found")
	}
	return clicks, nil
}

func (db *InMemoryDatabase) GetTenantByAPIKey(ctx context.Context, apiKey string) (*models.Tenant, error) {
	tenant, ok := db.tenants[apiKey]
	if !ok {
		return nil, errors.New("Invalid API key")
	}
	return &tenant, nil
}

func (db *InMemoryDatabase) Close() error {
	return nil
}
