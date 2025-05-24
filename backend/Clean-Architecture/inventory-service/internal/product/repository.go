package product

import (
	"inventory-service/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *Repository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r *Repository) Update(id uint, updates map[string]interface{}) error {
	return r.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updates).Error
}
func (r *Repository) Delete(id uint) error {
	return r.DB.Delete(&models.Product{}, id).Error
}
func (r *Repository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, id).Error
	return &product, err
}
