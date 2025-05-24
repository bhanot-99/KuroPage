package models

type Order struct {
    ID         uint    `json:"id" gorm:"primaryKey"`
    UserID     uint    `json:"user_id"`       // пользователь, сделавший заказ
    ProductID  uint    `json:"product_id"`    // продукт
    Quantity   int     `json:"quantity"`
    TotalPrice float64 `json:"total_price"`
    Status     string  `json:"status"` // например: pending, completed, cancelled
}
