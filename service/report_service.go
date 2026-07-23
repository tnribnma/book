package service

import (
	"context"
<<<<<<< HEAD
	"fmt"
	"book-management/models"
	"book-management/repository"
)

type ReportService interface {
	GetDashboard(ctx context.Context) (*models.DashboardStats, error)
	GetPopularBooks(ctx context.Context, limit int) ([]models.Book, error)
	GetSystemSummary(ctx context.Context) (*models.SystemSummary, error)
}

type reportService struct {
	reportRepo    repository.ReportRepository
	bookRepo      repository.BookRepository
	borrowingRepo repository.BorrowingRepository
}

func NewReportService(
	reportRepo repository.ReportRepository,
	bookRepo repository.BookRepository,
	borrowingRepo repository.BorrowingRepository,
) ReportService {
	return &reportService{
		reportRepo:    reportRepo,
		bookRepo:      bookRepo,
		borrowingRepo: borrowingRepo,
	}
}

func (s *reportService) GetDashboard(ctx context.Context) (*models.DashboardStats, error) {
	stats, err := s.reportRepo.GetDashboardStats(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get dashboard stats: %w", err)
	}
	return stats, nil
}

func (s *reportService) GetPopularBooks(ctx context.Context, limit int) ([]models.Book, error) {
	if limit <= 0 || limit > 20 {
		limit = 10 
	}

	books, err := s.reportRepo.GetPopularBooks(ctx, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get popular books: %w", err)
	}

	return books, nil
}

func (s *reportService) GetSystemSummary(ctx context.Context) (*models.SystemSummary, error) {
	summary := &models.SystemSummary{}

	totalBooks, err := s.reportRepo.GetTotalBooksCount(ctx)
	if err != nil {
		return nil, err
	}
	summary.TotalBooks = totalBooks

	activeBorrowings, err := s.reportRepo.GetActiveBorrowingsCount(ctx)
	if err != nil {
		return nil, err
	}
	summary.ActiveBorrowings = activeBorrowings

	overdue, err := s.reportRepo.GetOverdueCount(ctx)
	if err != nil {
		return nil, err
	}
	summary.OverdueBorrowings = overdue

	totalUsers, err := s.reportRepo.GetTotalUsersCount(ctx)
	if err != nil {
		return nil, err
	}
	summary.TotalUsers = totalUsers

	dashboard, err := s.reportRepo.GetDashboardStats(ctx)
	if err == nil {
		summary.AvailableBooks = dashboard.AvailableBooks
	}

	return summary, nil
=======
	"database/sql"

	"book-management/models"
)

type ReportService struct {
	db *sql.DB 
}

func NewReportService(db *sql.DB) *ReportService {
	return &ReportService{db: db}
}

func (s *ReportService) GetDashboardStats(ctx context.Context) (models.Report, error) {
	return models.Report{}, nil
}

func (s *ReportService) GetOverdueBooks(ctx context.Context) ([]models.Borrowing, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT b.id, b.book_id, bk.title, b.user_id, u.email,
			b.issue_date, b.due_date, b.return_date, b.status, b.fine_amount
		FROM borrowings b
		JOIN books bk ON bk.id = b.book_id
		JOIN users u ON u.id = b.user_id
		WHERE b.due_date < NOW() AND b.status = 'borrowed'
		ORDER BY b.due_date`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var overdue []models.Borrowing
	for rows.Next() {
		var b models.Borrowing
		if err := rows.Scan(&b.ID, &b.BookID, &b.BookTitle, &b.UserID, &b.UserEmail,
			&b.IssueDate, &b.DueDate, &b.ReturnDate, &b.Status, &b.FineAmount); err != nil {
			return nil, err
		}
		overdue = append(overdue, b)
	}
	return overdue, rows.Err()
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
}