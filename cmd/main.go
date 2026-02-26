package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/syaafiudinm/go-starter-kit/config"
	"github.com/syaafiudinm/go-starter-kit/internal/model"
	"github.com/syaafiudinm/go-starter-kit/routes"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	config.InitDatabase()

	// Run auto-migration
	config.AutoMigrate(
		&model.User{},
	)

	// Set Gin mode
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.Setup(router, config.DB)

	// Start server
	addr := ":" + cfg.App.Port
	log.Printf("%s is running on port %s in %s mode", cfg.App.Name, cfg.App.Port, cfg.App.Env)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
