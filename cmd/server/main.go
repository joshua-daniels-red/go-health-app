package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go-health-app/internal/config"
	"go-health-app/internal/handler"
	"go-health-app/internal/repository"
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

	// ----------------------------
	// Wire services and handlers
	// ----------------------------

	// Health
	healthService := service.NewHealthService()
	healthHandler := handler.NewHealthHandler(healthService)
	processHandler := handler.NewProcessHandler("python3", "python_app/main.py")


	// Data
	movieRepo, err := repository.NewMovieRepository("data/movies.json")
	if err != nil {
		log.Fatalf("failed to load movies: %v", err)
	}
	dataService := service.NewDataService(movieRepo)
	dataHandler := handler.NewDataHandler(dataService)

	// ----------------------------
	// Register routes
	// ----------------------------
	router.GET("/health", healthHandler.Check)
	router.GET("/data", dataHandler.GetData)
	router.GET("/process", processHandler.RunProcess)

	// Start server
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
