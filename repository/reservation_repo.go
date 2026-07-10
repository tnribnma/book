package repository

import (
	"context"
	"database/sql"
	"book-management/models"
)

type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

// Add Create, List, etc.