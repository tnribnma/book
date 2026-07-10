package repository

import (
	"context"
	"database/sql"
	"book-management/models"
)

type BorrowingRepository struct {
	db *sql.DB
}

func NewBorrowingRepository(db *sql.DB) *BorrowingRepository {
	return &BorrowingRepository{db: db}
}

func (r *BorrowingRepository) Create(ctx context.Context, borrowing models.Borrowing) (models.Borrowing, error) {
	var b models.Borrowing
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO borrowings (book_id, user_id, due_date)
		VALUES ($1, $2, $3) RETURNING id, book_id, user_id, issue_date, due_date, status`,
		borrowing.BookID, borrowing.UserID, borrowing.DueDate).Scan(
		&b.ID, &b.BookID, &b.UserID, &b.IssueDate, &b.DueDate, &b.Status)
	return b, err
}

func (r *BorrowingRepository) GetByID(ctx context.Context, id int64) (models.Borrowing, error) {
	// implement
	return models.Borrowing{}, nil
}

func (r *BorrowingRepository) ReturnBook(ctx context.Context, borrowingID int64, fine float64) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE borrowings SET return_date = NOW(), status = 'returned', fine_amount = $1 
		WHERE id = $2`, fine, borrowingID)
	return err
}

// Add methods: GetUserBorrowings, GetOverdue, etc.