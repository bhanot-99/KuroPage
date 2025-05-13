package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bhanot-99/KuroPage/Backend/manga-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type MangaRepository interface {
	CreateManga(ctx context.Context, manga *model.Manga) error
	UpdateManga(ctx context.Context, manga *model.Manga) error
	GetMangaByID(ctx context.Context, id string) (*model.Manga, error)
	ListMangas(ctx context.Context, page, limit int) ([]*model.Manga, int, error)
	UpdateStock(ctx context.Context, id string, quantity int) error
}

type mangaRepository struct {
	db *sqlx.DB
}

func NewMangaRepository(db *sqlx.DB) MangaRepository {
	return &mangaRepository{db: db}
}

func (r *mangaRepository) CreateManga(ctx context.Context, manga *model.Manga) error {
	query := `INSERT INTO mangas (title, author, description, price, stock) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.db.QueryRowContext(ctx, query,
		manga.Title, manga.Author, manga.Description, manga.Price, manga.Stock).Scan(&manga.ID)
}

func (r *mangaRepository) UpdateManga(ctx context.Context, manga *model.Manga) error {
	query := `UPDATE mangas SET title=$1, author=$2, description=$3, price=$4, stock=$5, updated_at=NOW() 
	          WHERE id=$6`
	_, err := r.db.ExecContext(ctx, query,
		manga.Title, manga.Author, manga.Description, manga.Price, manga.Stock, manga.ID)
	return err
}

func (r *mangaRepository) GetMangaByID(ctx context.Context, id string) (*model.Manga, error) {
	var manga model.Manga
	err := r.db.GetContext(ctx, &manga, "SELECT * FROM mangas WHERE id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &manga, nil
}

func (r *mangaRepository) ListMangas(ctx context.Context, page, limit int) ([]*model.Manga, int, error) {
	var mangas []*model.Manga
	offset := (page - 1) * limit

	err := r.db.SelectContext(ctx, &mangas,
		"SELECT * FROM mangas ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var total int
	err = r.db.GetContext(ctx, &total, "SELECT COUNT(*) FROM mangas")
	if err != nil {
		return nil, 0, err
	}

	return mangas, total, nil
}

func (r *mangaRepository) UpdateStock(ctx context.Context, id string, quantity int) error {
	query := `UPDATE mangas SET stock=stock-$1, updated_at=NOW() WHERE id=$2 AND stock>=$1`
	result, err := r.db.ExecContext(ctx, query, quantity, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("not enough stock or manga not found")
	}

	return nil
}
