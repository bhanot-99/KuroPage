package product

import (
	"inventory-service/internal/inventory/cache"
	"inventory-service/models"
)

type Service struct {
	Repo  *Repository
	Cache *cache.Cache
}

func (s *Service) CreateProduct(p *models.Product) error {
	// Сохраняем в БД
	if err := s.Repo.Create(p); err != nil {
		return err
	}
	// Добавляем в кэш
	s.Cache.Add(*p)
	return nil
}

func (s *Service) GetProducts() ([]models.Product, error) {
	// Возвращаем из кэша
	return s.Cache.GetAll(), nil
}

func (s *Service) GetProductByID(id uint) (*models.Product, error) {
	if item, found := s.Cache.GetByID(uint(id)); found {
		return &item, nil
	}
	return s.Repo.GetByID(id)
}

func (s *Service) UpdateProduct(id uint, updates map[string]interface{}) error {
	return s.Repo.Update(id, updates)
}

func (s *Service) DeleteProduct(id uint) error {
	return s.Repo.Delete(id)
}
