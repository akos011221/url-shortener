package storage

import (
	"context"
	"errors"
)

type Database interface {
	SaveURL(ctx context.Context, shortCode, longURL string) error
	GetURL(ctx context.Context, shortCode string) (string, error)
	Close() error
}

type InMemoryDatabase struct {
	urls map[string]string
}

func NewDatabase(databaseURL string) (Database, error) {
	// TODO: Connect to a real database (e.g., PostgreSQL, Redis)
	return &InMemoryDatabase{urls: make(map[string]string)}, nil
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

func (db *InMemoryDatabase) Close() error {
	return nil
}
