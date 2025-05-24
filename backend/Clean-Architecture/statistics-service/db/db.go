package db

import (
	"log"

	"github.com/elliaaan/statistics-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=kingston dbname=statistics_db port=1054 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Автомиграция таблицы Event
	if err := db.AutoMigrate(&models.Event{}); err != nil {
		log.Fatalf("failed to migrate event model: %v", err)
	}

	return db
}
