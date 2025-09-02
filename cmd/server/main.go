package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go-health-app/internal/config"
	"go-health-app/internal/handler"
	"go-health-app/internal/service"
)

func main() {
	// Load configuration (from env, flags, defaults)
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Create Gin router
	router := gin.Default()

	// Wire services
	healthService := service.NewHealthService()

	// Wire handlers with services
	healthHandler := handler.NewHealthHandler(healthService)

	// Register routes
	router.GET("/health", healthHandler.Check)

	// Start server
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
