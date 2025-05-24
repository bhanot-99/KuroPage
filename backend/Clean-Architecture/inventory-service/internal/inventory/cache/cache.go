package cache

import (
	"log"
	"sync"
	"time"

	"inventory-service/models"

	"gorm.io/gorm"
)

type Cache struct {
	data map[uint]models.Product // 🔁 uint вместо uint64
	mu   sync.RWMutex
	db   *gorm.DB
}

// NewCache создаёт кэш и запускает автообновление
func NewCache(db *gorm.DB) *Cache {
	c := &Cache{
		data: make(map[uint]models.Product),
		db:   db,
	}
	c.LoadFromDB()
	go c.startAutoRefresh()
	return c
}

// LoadFromDB заполняет кэш из базы данных
func (c *Cache) LoadFromDB() {
	var items []models.Product
	if err := c.db.Find(&items).Error; err != nil {
		log.Println("Ошибка загрузки кэша из БД:", err)
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = make(map[uint]models.Product) // сброс старых данных
	for _, item := range items {
		c.data[item.ID] = item
	}
	log.Printf("[Cache] Загружено %d элементов из базы", len(c.data))
}

// startAutoRefresh — обновляет кэш каждые 12 часов
func (c *Cache) startAutoRefresh() {
	ticker := time.NewTicker(12 * time.Hour)
	for range ticker.C {
		log.Println("[Cache] Плановое обновление...")
		c.LoadFromDB()
	}
}

// GetAll — возвращает все элементы из кэша
func (c *Cache) GetAll() []models.Product {
	c.mu.RLock()
	defer c.mu.RUnlock()

	items := make([]models.Product, 0, len(c.data))
	for _, item := range c.data {
		items = append(items, item)
	}
	return items
}

// GetByID — возвращает продукт по ID
func (c *Cache) GetByID(id uint) (models.Product, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.data[id]
	return item, exists
}

// Add — добавляет продукт в кэш
func (c *Cache) Add(item models.Product) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[item.ID] = item
}
