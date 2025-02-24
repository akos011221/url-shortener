package main

import (
	"log"
	"net/http"

	"github.com/akos011221/url-shortener/api"
	"github.com/akos011221/url-shortener/config"
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

	// Create HTTP server with routes
	router := api.NewRouter(shortenerService, analyticsService)

	// Start the server
	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
