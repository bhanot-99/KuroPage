package model

import "time"

type Order struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	TotalAmount float64   `db:"total_amount"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type OrderItem struct {
	ID        string    `db:"id"`
	OrderID   string    `db:"order_id"`
	MangaID   string    `db:"manga_id"`
	Quantity  int       `db:"quantity"`
	UnitPrice float64   `db:"unit_price"`
	CreatedAt time.Time `db:"created_at"`
}
