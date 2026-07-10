package service

import (
	"context"

	"book-management/models"
	"book-management/repository"
)

type ReportService struct {
	db *sql.DB // or use specific repos
}

func NewReportService(db *sql.DB) *ReportService {
	return &ReportService{db: db}
}

func (s *ReportService) GetDashboardStats(ctx context.Context) (models.Report, error) {
	// Run aggregate queries
	return models.Report{}, nil
}

func (s *ReportService) GetOverdueBooks(ctx context.Context) ([]models.Borrowing, error) {
	// Query borrowings where due_date < now and status = 'borrowed'
	return nil, nil
}