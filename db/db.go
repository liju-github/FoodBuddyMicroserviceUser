package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/liju-github/EcommerceUserService/configs"
	"github.com/liju-github/EcommerceUserService/models"
)

// Connect establishes a connection to the SQLite database using GORM and configures connection pool settings.
// Returns a GORM database instance or an error if the connection fails.
func Connect(cfg config.Config) (*gorm.DB, error) {
	dsn := "./db.sqlite3" // Database path; can be adjusted in the config if needed

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve SQL DB instance: %w", err)
	}

	// Configure database connection pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	// Auto-migrate database schema for all models
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, fmt.Errorf("auto-migration failed: %w", err)
	}

	log.Println("Connected to SQLite database and schema migrated")
	return db, nil
}

// Close terminates the SQLite database connection safely.
func Close(db *gorm.DB) {
	if db == nil {
		log.Println("No active database connection to close.")
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Failed to retrieve SQL DB instance for closure: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("Failed to close database connection: %v", err)
		return
	}

	log.Println("Database connection closed")
}
