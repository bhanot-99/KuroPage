package statistics

import (
	"encoding/json"
	"log"

	"github.com/elliaaan/statistics-service/models"

	"time"

	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

// Структура события
type OrderCreatedEvent struct {
	OrderID    uint    `json:"order_id"`
	UserID     uint    `json:"user_id"`
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}

func SubscribeToOrderCreated(nc *nats.Conn, db *gorm.DB) {
	_, err := nc.Subscribe("order.created", func(msg *nats.Msg) {
		var event OrderCreatedEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Ошибка при разборе события: %v", err)
			return
		}

		// Логируем
		log.Printf("Получено событие: %+v", event)

		// Сохраняем в БД
		eventRecord := models.Event{
			Type:      "order.created",
			Data:      string(msg.Data),
			CreatedAt: time.Now(),
		}

		if err := db.Create(&eventRecord).Error; err != nil {
			log.Printf("Ошибка при сохранении события: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("Ошибка подписки на NATS: %v", err)
	}
}
