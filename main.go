package main

import (
	"log"
	"net/http"

	"github.com/akos011221/url-shortener/api"
	"github.com/akos011221/url-shortener/config"
	"github.com/akos011221/url-shortener/service"
	"github.com/akos011221/url-shortener/storage"
	"github.com/akos011221/url-shortener/utils"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	utils.InitLogger(cfg.Env)

	// Initialize database
	db, err := storage.NewDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create services
	shortenerService := service.NewShortener(db)
	analyticsService := service.NewAnalytics(db)

	// Create handlers
	handlers := api.Handlers{
		Shortener: shortenerService,
		Analytics: analyticsService,
	}

	// Create router for public routes
	router := http.NewServeMux()

	// Public routes (no API key required)
	router.HandleFunc("GET /{shortCode}", handlers.Redirect) // Redirect short URLs

	// Create router for protected routes
	protectedRouter := http.NewServeMux()

	// Protected routes (no API key required)
	protectedRouter.HandleFunc("POST /shorten", handlers.CreateShortURL)            // Create short URL
	protectedRouter.HandleFunc("GET /analytics/{shortCode}", handlers.GetAnalytics) // Get analytics

	// Wrap the protected router with middleware
	protectedHandler := api.LoggingMiddleware(protectedRouter)
	protectedHandler = api.RateLimitMiddleware(protectedHandler)
	protectedHandler = api.AuthMiddleware(db, protectedHandler)

	// Mount the protected router under a prefix
	router.Handle("/api/", http.StripPrefix("/api", protectedHandler))

	// Wrap the public router with logging and rate limit middleware
	handler := api.LoggingMiddleware(router)
	handler = api.RateLimitMiddleware(handler)

	// Start the server
	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler,
	}

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
