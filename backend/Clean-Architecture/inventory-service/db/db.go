package db

import (
	"inventory-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// 1️ Connection string (DSN) to PostgreSQL
	dsn := "host=localhost user=postgres password=kingston dbname=inventory_db port=1054 sslmode=disable"

	// 2️ Connect to PostgreSQL using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// If connection fails → stop the program
		panic("failed to connect to PostgreSQL database")
	}

	db.Migrator().DropTable(&models.Product{}) // удалить старую таблицу
	db.AutoMigrate(&models.Product{})          // создать заново с новым autoIncrement

	// 4️ Return DB object to be used in Repository
	return db
}
