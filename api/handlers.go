package api

import (
	"encoding/json"
	"errors"
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
		utils.WriteError(w, http.StatusBadRequest, utils.ErrInvalidRequestBody)
		return
	}

	// Validate request
	if req.LongURL == "" {
		utils.WriteError(w, http.StatusBadRequest, utils.ErrLongURLRequired)
		return
	}

	// Extract tenant ID from context (set by AuthMiddleware)
	tenantID, ok := r.Context().Value("tenantID").(string)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, utils.ErrUnauthorizedAccess)
	}

	// Create short URL
	shortURL, err := h.Shortener.CreateShortURL(r.Context(), req.LongURL, tenantID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, utils.ErrFailedToCreateURL)
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
		utils.WriteError(w, http.StatusNotFound, utils.ErrURLNotFound)
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

	// Extract tenant ID from context (set by AuthMiddleware)
	tenantID := r.Context().Value("tenantID").(string)

	// Get the short URL's tenant ID (if any)
	urlTenantID, err := h.Shortener.GetURLTenantID(r.Context(), shortCode)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, utils.ErrURLNotFound)
		return
	}

	// Ensure the tenant is authorized to access the analytics
	if urlTenantID != "" && urlTenantID != tenantID {
		utils.WriteError(w, http.StatusForbidden, utils.ErrUnauthorizedAccess)
		return
	}

	// Get analytics data
	analytics, err := h.Analytics.GetAnalytics(r.Context(), shortCode)
	if err != nil {
		// If no clicks are found, return a 200 OK response with zero clicks
		if errors.Is(err, utils.ErrNoClicksFound) {
			utils.WriteJSON(w, http.StatusOK, models.GetAnalyticsResponse{
				ShortCode: shortCode,
				Clicks:    0,
				Details:   []models.Click{},
			})
			return
		}

		// For other errors, return a 500 Internal Server Error
		utils.WriteError(w, http.StatusInternalServerError, utils.ErrFailedToRetrieveData)
		return
	}

	// Return response
	utils.WriteJSON(w, http.StatusOK, analytics)
}
