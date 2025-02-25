package api

import (
	"encoding/json"
	"net/http"

	"github.com/akos011221/url-shortener/models"
	"github.com/akos011221/url-shortener/service"
	"github.com/akos011221/url-shortener/utils"
)

type Handlers struct {
	Shortener *service.Shortener
	Analytics *service.Analytics
}

// CreateShortURL handles requests to create a short URL.
func (h *Handlers) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req models.CreateShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate request
	if req.LongURL == "" {
		utils.WriteError(w, http.StatusBadRequest, "Long URL is required")
		return
	}

	// Extract API key from request headers
	apiKey := r.Header.Get("x-api-key")
	var tenantID string

	// If API key is provided, validate it and get the tenant ID
	if apiKey != "" {
		tenant, err := h.Shortener.GetTenantByAPIKey(r.Context(), apiKey)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, "Invalid API key")
			return
		}
		tenantID = tenant.ID
	}

	// Create short URL
	shortURL, err := h.Shortener.CreateShortURL(r.Context(), req.LongURL, tenantID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to create short URL")
		return
	}

	// Return response
	utils.WriteJSON(w, http.StatusCreated, models.CreateShortURLResponse{ShortURL: shortURL})
}

// Redirect handles requests to redirect to the original URL.
func (h *Handlers) Redirect(w http.ResponseWriter, r *http.Request) {
	// Extract short code from URL path
	shortCode := r.URL.Path[1:]

	// Get long URL
	longURL, err := h.Shortener.GetLongURL(r.Context(), shortCode)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Short URL not found")
		return
	}

	// Log analytics data
	go h.Analytics.LogClick(r.Context(), shortCode, r.RemoteAddr, r.UserAgent())

	// Redirect to the long URL
	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

// GetAnalytics handles requests to retrieve analytics for a short URL.
func (h *Handlers) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	// Extract short code from URL path
	shortCode := r.PathValue("shortCode")

	// Extract API key from request headers
	apiKey := r.Header.Get("x-api-key")
	if apiKey == "" {
		utils.WriteError(w, http.StatusUnauthorized, "API key required")
		return
	}

	// Validate API key and get the tenant
	tenant, err := h.Shortener.GetTenantByAPIKey(r.Context(), apiKey)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid API key")
		return
	}

	// Get the short URL's tenant ID (if any)
	urlTenantID, err := h.Shortener.GetURLTenantID(r.Context(), shortCode)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Short URL not found")
		return
	}

	// Ensure the tenant is authorized to access the analytics
	if urlTenantID != "" && urlTenantID != tenant.ID {
		utils.WriteError(w, http.StatusForbidden, "You are not authorized to access the short URL's analytics")
		return
	}

	// Get analytics data
	analytics, err := h.Analytics.GetAnalytics(r.Context(), shortCode)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to retrieve analytics")
		return
	}

	// Return response
	utils.WriteJSON(w, http.StatusOK, analytics)
}
