package storage

import (
	"context"

	"github.com/akos011221/url-shortener/models"
	"github.com/akos011221/url-shortener/utils"
)

type Database interface {
	SaveURL(ctx context.Context, shortCode, longURL, tenantID string) error
	GetURL(ctx context.Context, shortCode string) (string, error)
	SaveClick(ctx context.Context, click models.Click) error
	GetClicks(ctx context.Context, shortCode string) ([]models.Click, error)
	GetTenantByAPIKey(ctx context.Context, apiKey string) (*models.Tenant, error)
	GetURLTenantID(ctx context.Context, shortCode string) (string, error)
	Close() error
}

type InMemoryDatabase struct {
	urls       map[string]string         // shortCode -> longURL
	clicks     map[string][]models.Click // shortCode -> []Click
	tenants    map[string]models.Tenant  // apiKey -> Tenant
	urlTenants map[string]string	     // shortCode -> tenantID
}

func NewDatabase(databaseURL string) (Database, error) {
	// Initialize in-memory data
	tenants := map[string]models.Tenant{
		"api-key-123": {ID: "1", Name: "Tenant A", APIKey: "api-key-123"},
	}

	// TODO: Connect to a real database (e.g., PostgreSQL, Redis)
	return &InMemoryDatabase{
		urls:       make(map[string]string),
		clicks:     make(map[string][]models.Click),
		tenants:    tenants,
		urlTenants: make(map[string]string),
	}, nil
}

func (db *InMemoryDatabase) SaveURL(ctx context.Context, shortCode, longURL, tenantID string) error {
	db.urls[shortCode] = longURL
	if tenantID != "" {
		db.urlTenants[shortCode] = tenantID
	}
	return nil
}

func (db *InMemoryDatabase) GetURL(ctx context.Context, shortCode string) (string, error) {
	longURL, ok := db.urls[shortCode]
	if !ok {
		return "", utils.ErrURLNotFound
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
		return nil, utils.ErrNoClicksFound
	}
	return clicks, nil
}

func (db *InMemoryDatabase) GetTenantByAPIKey(ctx context.Context, apiKey string) (*models.Tenant, error) {
	tenant, ok := db.tenants[apiKey]
	if !ok {
		return nil, utils.ErrInvalidAPIKey
	}
	return &tenant, nil
}

func (db *InMemoryDatabase) GetURLTenantID(ctx context.Context, shortCode string) (string, error) {
	if _, ok := db.urls[shortCode]; !ok {
		return "", utils.ErrURLNotFound
	}
	return db.urlTenants[shortCode], nil
}

func (db *InMemoryDatabase) Close() error {
	return nil
}
