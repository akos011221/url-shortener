package service

import (
	"context"
	"time"

	"github.com/akos011221/url-shortener/models"
	"github.com/akos011221/url-shortener/storage"
)

type Analytics struct {
	db storage.Database
}

func NewAnalytics(db storage.Database) *Analytics {
	return &Analytics{db: db}
}

// LogClick records a click on a short URL.
func (a *Analytics) LogClick(ctx context.Context, shortCode, ipAddress, userAgent string) error {
	// Create a Click object
	click := models.Click{
		ShortCode: shortCode,
		IPAddress: ipAddress,
		UserAgent: userAgent,
		TimeStamp: time.Now(),
	}

	// Save the click to the database
	if err := a.db.SaveClick(ctx, click); err != nil {
		return err
	}

	return nil
}

// GetAnalytics retrieves analytics data for a short URL.
func (a *Analytics) GetAnalytics(ctx context.Context, shortCode string) (*models.GetAnalyticsResponse, error) {
	// Get total clicks
	clicks, err := a.db.GetClicks(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	// Prepare response
	response := &models.GetAnalyticsResponse{
		ShortCode: shortCode,
		Clicks:    len(clicks),
		Details:   clicks,
	}

	return response, nil
}
