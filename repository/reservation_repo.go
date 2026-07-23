package repository

import (
	"context"
	"database/sql"
<<<<<<< HEAD
	"fmt"
	"book-management/models"
)

type ReservationRepository interface {
	Create(ctx context.Context, reservation *models.Reservation) error
	GetByID(ctx context.Context, id int64) (*models.Reservation, error)
	GetUserReservations(ctx context.Context, userID int64) ([]models.Reservation, error)
	Cancel(ctx context.Context, id int64, userID int64) error
	Fulfill(ctx context.Context, reservationID, bookID int64) error
	HasActiveReservation(ctx context.Context, bookID, userID int64) (bool, error)
}

type reservationRepo struct {
	db *sql.DB
}

func (r *reservationRepo) HasActiveReservation(ctx context.Context, bookID, userID int64) (bool, error) {
	var count int
	err := r.db.QueryRowContext(ctx, `
		SELECT COUNT(*) FROM reservations 
		WHERE book_id = $1 AND user_id = $2 AND status = 'pending'`, 
		bookID, userID).Scan(&count)
	return count > 0, err
}

func NewReservationRepository(db *sql.DB) ReservationRepository {
	return &reservationRepo{db: db}
}

func (r *reservationRepo) Create(ctx context.Context, reservation *models.Reservation) error {
	query := `
		INSERT INTO reservations (book_id, user_id, status)
		VALUES ($1, $2, 'pending')
		RETURNING id, reservation_date`

	err := r.db.QueryRowContext(ctx, query, 
		reservation.BookID, 
		reservation.UserID,
	).Scan(&reservation.ID, &reservation.ReservationDate)

	if err != nil {
		return fmt.Errorf("failed to create reservation: %w", err)
	}

	_, err = r.db.ExecContext(ctx, `
		UPDATE books 
		SET status = 'reserved' 
		WHERE id = $1`, reservation.BookID)

	if err != nil {
		return fmt.Errorf("failed to update book status: %w", err)
	}

	return nil
}

func (r *reservationRepo) GetByID(ctx context.Context, id int64) (*models.Reservation, error) {
	query := `
		SELECT r.id, r.book_id, bk.title as book_title, r.user_id, u.full_name as user_name,
		       r.reservation_date, r.status, r.created_at
		FROM reservations r
		JOIN books bk ON r.book_id = bk.id
		JOIN users u ON r.user_id = u.id
		WHERE r.id = $1`

	var res models.Reservation
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&res.ID, &res.BookID, &res.BookTitle, &res.UserID, &res.UserName,
		&res.ReservationDate, &res.Status, &res.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("reservation not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get reservation: %w", err)
	}

	return &res, nil
}

func (r *reservationRepo) GetUserReservations(ctx context.Context, userID int64) ([]models.Reservation, error) {
	query := `
		SELECT r.id, r.book_id, bk.title as book_title, r.user_id, 
		       r.reservation_date, r.status, r.created_at
		FROM reservations r
		JOIN books bk ON r.book_id = bk.id
		WHERE r.user_id = $1
		ORDER BY r.reservation_date DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user reservations: %w", err)
=======
	"book-management/models"
)

type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) Create(ctx context.Context, reservation models.Reservation) (models.Reservation, error) {
	var created models.Reservation
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO reservations (user_id, book_id, status, reservation_date)
		VALUES ($1, $2, 'pending', NOW())
		RETURNING id, user_id, book_id, status, reservation_date`,
		reservation.UserID, reservation.BookID).Scan(
		&created.ID,
		&created.UserID,
		&created.BookID,
		&created.Status,
		&created.ReservationDate,
	)
	return created, err
}

func (r *ReservationRepository) List(ctx context.Context, userID int64) ([]models.Reservation, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, book_id, status, reservation_date 
		FROM reservations 
		WHERE user_id = $1 
		ORDER BY reservation_date DESC`, userID)
	if err != nil {
		return nil, err
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
	}
	defer rows.Close()

	var reservations []models.Reservation
	for rows.Next() {
		var res models.Reservation
<<<<<<< HEAD
		if err := rows.Scan(
			&res.ID, &res.BookID, &res.BookTitle, &res.UserID,
			&res.ReservationDate, &res.Status, &res.CreatedAt,
		); err != nil {
=======
		err := rows.Scan(
			&res.ID,
			&res.UserID,
			&res.BookID,
			&res.Status,
			&res.ReservationDate,
		)
		if err != nil {
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
			return nil, err
		}
		reservations = append(reservations, res)
	}
<<<<<<< HEAD

	return reservations, nil
}

func (r *reservationRepo) Cancel(ctx context.Context, id int64, userID int64) error {
	query := `
		UPDATE reservations 
		SET status = 'cancelled'
		WHERE id = $1 AND user_id = $2 AND status = 'pending'`

	result, err := r.db.ExecContext(ctx, query, id, userID)
	if err != nil {
		return fmt.Errorf("failed to cancel reservation: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("reservation not found or already processed")
	}

	return nil
}

func (r *reservationRepo) Fulfill(ctx context.Context, reservationID, bookID int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `
		UPDATE reservations 
		SET status = 'fulfilled' 
		WHERE id = $1`, reservationID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		UPDATE books 
		SET status = 'available' 
		WHERE id = $1`, bookID)
	if err != nil {
		return err
	}

	return tx.Commit()
=======
	return reservations, nil
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
}