package order

import (
	"order-service/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *Repository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Find(&orders).Error
	return orders, err
}

func (r *Repository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.DB.First(&order, id).Error
	return &order, err
}

func (r *Repository) UpdateStatus(id uint, status string) error {
	return r.DB.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
}
func (r *Repository) GetByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
func (r *Repository) Update(id uint, updates map[string]interface{}) error {
	return r.DB.Model(&models.Order{}).Where("id = ?", id).Updates(updates).Error
}

func (r *Repository) Delete(id uint) error {
	return r.DB.Delete(&models.Order{}, id).Error
}
