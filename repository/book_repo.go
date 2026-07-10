package repository

import (
	"context"
	"database/sql"
	"fmt"
	"book-management/models"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(ctx context.Context, book models.Book) (models.Book, error) {
	var created models.Book
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO books (title, author, isbn, category_id, publisher, edition, 
			published_year, quantity, available_copies, shelf, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $8, $9, 'available')
		RETURNING id, title, author, isbn, category_id, publisher, edition, 
			published_year, quantity, available_copies, shelf, status, created_at`,
		book.Title, book.Author, book.ISBN, book.CategoryID, book.Publisher,
		book.Edition, book.PublishedYear, book.Quantity, book.Shelf).Scan(
		&created.ID, &created.Title, &created.Author, &created.ISBN, &created.CategoryID,
		&created.Publisher, &created.Edition, &created.PublishedYear, &created.Quantity,
		&created.AvailableCopies, &created.Shelf, &created.Status, &created.CreatedAt)
	return created, err
}

func (r *BookRepository) List(ctx context.Context, filter models.BookFilter, limit, offset int) ([]models.Book, int, error) {
	// Implement with search, filter, pagination (basic version shown)
	query := `SELECT id, title, author, isbn, category_id, publisher, edition, 
		published_year, quantity, available_copies, shelf, status, created_at 
		FROM books WHERE 1=1`
	// Add WHERE conditions dynamically for search etc.
	// ... (full implementation can be expanded)
	return nil, 0, nil // placeholder - expand as needed
}

func (r *BookRepository) GetByID(ctx context.Context, id int64) (models.Book, error) {
	var book models.Book
	err := r.db.QueryRowContext(ctx, `SELECT * FROM books WHERE id = $1`, id).Scan(
		&book.ID, &book.Title, &book.Author, &book.ISBN, &book.CategoryID,
		&book.Publisher, &book.Edition, &book.PublishedYear, &book.Quantity,
		&book.AvailableCopies, &book.Shelf, &book.Status, &book.CreatedAt)
	return book, err
}

func (r *BookRepository) UpdateAvailability(ctx context.Context, bookID int64, change int) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE books SET available_copies = available_copies + $1 
		WHERE id = $2`, change, bookID)
	return err
}