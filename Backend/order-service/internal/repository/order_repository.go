package repository

import (
	"context"

	"github.com/bhanot-99/KuroPage/Backend/order-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *model.Order, items []*model.OrderItem) error
}

type orderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order *model.Order, items []*model.OrderItem) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	// Create order
	orderQuery := `INSERT INTO orders (user_id, total_amount) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRowxContext(ctx, orderQuery, order.UserID, order.TotalAmount).Scan(&order.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Create order items
	itemQuery := `INSERT INTO order_items (order_id, manga_id, quantity, unit_price) 
	              VALUES ($1, $2, $3, $4)`
	for _, item := range items {
		_, err = tx.ExecContext(ctx, itemQuery,
			order.ID, item.MangaID, item.Quantity, item.UnitPrice)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
