package statistics

import (
	"github.com/elliaaan/statistics-service/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAllEvents() ([]models.Event, error) {
	var events []models.Event
	err := r.DB.Order("created_at desc").Find(&events).Error
	return events, err
}
