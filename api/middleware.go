package api

import (
	"log"
	"net/http"
	"time"
	"context"

	"github.com/akos011221/url-shortener/utils"
	"github.com/akos011221/url-shortener/storage"
)

// LoggingMiddleware logs incoming requests.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// AuthMiddleware authenticates tenants using API keys.
func AuthMiddleware(db storage.Database, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if apiKey == "" {
			utils.WriteError(w, http.StatusUnauthorized, utils.ErrAPIKeyRequired)
			return
		}

		// Validate API key and get the tenant
		tenant, err := db.GetTenantByAPIKey(r.Context(), apiKey)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, utils.ErrInvalidAPIKey)
			return
		}

		// Log the tenant for debugging
		log.Printf("Authenticated tenant: %s", tenant.Name)
		
		// Add tenant ID to the request context
		ctx := context.WithValue(r.Context(), "tenantID", tenant.ID)

		// Call the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RateLimitMiddleware enforces rate limits pers tenant.
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		// TODO: implement rate limiting logic
		next.ServeHTTP(w, r)
	})
}
