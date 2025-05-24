package order

import (
	"encoding/json"
	"order-service/models"

	"github.com/nats-io/nats.go"
)

type Service struct {
	Repo  *Repository
	Cache *Cache
	NATS  *nats.Conn
}

func (s *Service) CreateOrder(order *models.Order) error {

	if err := s.Repo.Create(order); err != nil {
		return err
	}

	s.Cache.Add(*order)

	event := map[string]interface{}{
		"order_id":    order.ID,
		"user_id":     order.UserID,
		"product_id":  order.ProductID,
		"quantity":    order.Quantity,
		"total_price": order.TotalPrice,
		"status":      order.Status,
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if err := s.NATS.Publish("order.created", data); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetOrderByID(id uint) (*models.Order, error) {
	if order, ok := s.Cache.GetByID(id); ok {
		return &order, nil
	}
	return s.Repo.GetByID(id)
}

func (s *Service) ListOrders() ([]models.Order, error) {
	return s.Cache.GetAll(), nil
}

func (s *Service) UpdateOrder(order *models.Order) error {
	updates := map[string]interface{}{
		"user_id":     order.UserID,
		"product_id":  order.ProductID,
		"quantity":    order.Quantity,
		"total_price": order.TotalPrice,
		"status":      order.Status,
	}
	return s.Repo.Update(order.ID, updates)
}

func (s *Service) DeleteOrder(id uint) error {
	return s.Repo.Delete(id)
}
