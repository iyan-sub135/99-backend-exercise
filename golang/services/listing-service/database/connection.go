package database

import (
	"listing-service/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database/listings.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Listing{}); err != nil {
		return nil, err
	}

	return db, nil
}
