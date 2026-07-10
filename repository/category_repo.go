package repository

import (
	"context"
	"database/sql"
	"book-management/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, cat models.Category) (models.Category, error) {
	// INSERT logic
	return cat, nil
}

func (r *CategoryRepository) List(ctx context.Context) ([]models.Category, error) {
	// SELECT all
	return nil, nil
}