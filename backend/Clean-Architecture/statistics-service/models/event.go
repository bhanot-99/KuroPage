package models

import "time"

type Event struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string // например: "order.created", "product.updated"
	Data      string // JSON или сериализованное содержимое события
	CreatedAt time.Time
}
