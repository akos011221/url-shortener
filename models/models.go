package models

import (
	"time"
)

// Tenant represents a user or organization using the URL shortener.
type Tenant struct {
	ID	string	`json:"id"`
	Name	string `json:"name"`
	APIKey string	`json:"apiKey"`
}

// URL represents a shortened URL and its metadata.
type URL struct {
	ShortCode	string	`json:"shortCode"`
	LongURL		string	`json:"longURL"`
	CreatedAt 	time.Time `json:"createdAt"`
	TenantID	string	`json:"tenantId"`
}

// Click represents a click on a short URL.
type Click struct {
	ShortCode	string	`json:"shortCode"`
	IPAddress	string	`json:"ipAddress"`
	UserAgent	string	`json:"userAgent"`
	TimeStamp	time.Time `json:"timestamp"`
}

// CreateShortURLRequest represents a request body for creating a short URL.
type CreateShortURLRequest struct {
	LongURL string `json:"longUrl"`
}

type CreateShortURLResponse struct {
	ShortURL string `json:"shortUrl"`
}

// GetAnalyticsResponse represents the response body for analytics data.
type GetAnalyticsResponse struct {
	ShortCode string `json:"shortCode"`
	Clicks	int	`json:"clicks"`
	Details	[]Click	`json:"details"`
}
