package main

import (
	"lion_parcel/internal/config"
	"lion_parcel/internal/handlers"
	"lion_parcel/internal/models"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger" // Import echo-swagger
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPass + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate models
	db.AutoMigrate(&models.Movie{}, &models.User{}, &models.Vote{})

	// Initialize Echo framework
	e := echo.New()

	// Setup routes
	handlers.SetupRoutes(e, db)

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := e.Start(":" + cfg.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
