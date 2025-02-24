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

	// Create router
	router := http.NewServeMux()

	// Register routes
	router.HandleFunc("POST /shorten", handlers.CreateShortURL)
	router.HandleFunc("GET /{shortCode}", handlers.Redirect)

	// Wrap the router with middleware
	handler := api.LoggingMiddleware(router)
	handler = api.AuthMiddleware(handler)
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
