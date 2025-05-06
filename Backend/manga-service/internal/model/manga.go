package model

import "time"

type Manga struct {
	ID          string    `db:"id"`
	Title       string    `db:"title"`
	Author      string    `db:"author"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Stock       int       `db:"stock"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
