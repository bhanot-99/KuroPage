package order

import (
	"log"
	"sync"
	"time"

	"order-service/models"

	"gorm.io/gorm"
)

type Cache struct {
	data map[uint]models.Order
	mu   sync.RWMutex
	db   *gorm.DB
}

func NewCache(db *gorm.DB) *Cache {
	c := &Cache{
		data: make(map[uint]models.Order),
		db:   db,
	}
	c.LoadFromDB()
	go c.startAutoRefresh()
	return c
}

func (c *Cache) LoadFromDB() {
	var orders []models.Order
	if err := c.db.Find(&orders).Error; err != nil {
		log.Println("Ошибка загрузки заказов в кэш:", err)
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = make(map[uint]models.Order)
	for _, order := range orders {
		c.data[order.ID] = order
	}
	log.Printf("[OrderCache] Загружено %d заказов из базы", len(c.data))
}

func (c *Cache) startAutoRefresh() {
	ticker := time.NewTicker(15 * time.Second)
	for range ticker.C {
		log.Println("[OrderCache] Плановое обновление...")
		c.LoadFromDB()
	}
}

func (c *Cache) GetAll() []models.Order {
	c.mu.RLock()
	defer c.mu.RUnlock()

	orders := make([]models.Order, 0, len(c.data))
	for _, order := range c.data {
		orders = append(orders, order)
	}
	return orders
}

func (c *Cache) GetByID(id uint) (models.Order, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	order, exists := c.data[id]
	return order, exists
}

func (c *Cache) Add(order models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[order.ID] = order
}
