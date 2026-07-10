package service

import (
	"context"
	"errors"
	"time"

	"book-management/models"
	"book-management/repository"
)

type BorrowingService struct {
	bookRepo      *repository.BookRepository
	borrowingRepo *repository.BorrowingRepository
}

func NewBorrowingService(db *sql.DB) *BorrowingService {
	return &BorrowingService{
		bookRepo:      repository.NewBookRepository(db),
		borrowingRepo: repository.NewBorrowingRepository(db),
	}
}

func (s *BorrowingService) IssueBook(ctx context.Context, req models.BorrowRequest, userID int64) (models.Borrowing, error) {
	book, err := s.bookRepo.GetByID(ctx, req.BookID)
	if err != nil {
		return models.Borrowing{}, errors.New("book not found")
	}

	if book.AvailableCopies <= 0 {
		return models.Borrowing{}, errors.New("no available copies")
	}

	dueDate := time.Now().AddDate(0, 0, req.DueDays)
	if dueDate.Before(time.Now()) {
		return models.Borrowing{}, errors.New("due date must be in future")
	}

	borrowing := models.Borrowing{
		BookID:  req.BookID,
		UserID:  userID,
		DueDate: dueDate,
	}

	b, err := s.borrowingRepo.Create(ctx, borrowing)
	if err != nil {
		return b, err
	}

	// Update book availability
	s.bookRepo.UpdateAvailability(ctx, req.BookID, -1)

	return b, nil
}

func (s *BorrowingService) ReturnBook(ctx context.Context, borrowingID int64) (float64, error) {
	fine := 0.0
	// TODO: Calculate overdue fine (e.g. 50 per day)

	err := s.borrowingRepo.ReturnBook(ctx, borrowingID, fine)
	if err != nil {
		return 0, err
	}

	// Increase available copies (need book_id)
	return fine, nil
}

func (s *BorrowingService) GetUserBorrowings(ctx context.Context, userID int64) ([]models.Borrowing, error) {
	// Implement in repository
	return nil, nil
}