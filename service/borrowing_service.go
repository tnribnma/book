package service

import (
	"context"
<<<<<<< HEAD
	"fmt"
=======
	"errors"
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
	"time"

	"book-management/models"
	"book-management/repository"
<<<<<<< HEAD
)

type BorrowingService interface {
	IssueBook(ctx context.Context, bookID, userID int64) error
	ReturnBook(ctx context.Context, bookID, userID int64) error
	GetMyBorrowings(ctx context.Context, userID int64) ([]models.Borrowing, error)
	GetOverdueBorrowings(ctx context.Context) ([]models.Borrowing, error)
}

type borrowingService struct {
	borrowingRepo repository.BorrowingRepository
	bookRepo      repository.BookRepository
}

func NewBorrowingService(borrowingRepo repository.BorrowingRepository, bookRepo repository.BookRepository) BorrowingService {
	return &borrowingService{
		borrowingRepo: borrowingRepo,
		bookRepo:      bookRepo,
	}
}

func (s *borrowingService) IssueBook(ctx context.Context, bookID, userID int64) error {
	book, err := s.bookRepo.GetByID(ctx, bookID)
	if err != nil {
		return fmt.Errorf("book not found")
	}

	if book.AvailableCopies <= 0 {
		return fmt.Errorf("book is currently not available")
	}

	// Check if user already borrowed this book
	hasActive, err := s.borrowingRepo.HasActiveBorrowing(ctx, bookID, userID)
	if err != nil {
		return fmt.Errorf("failed to check borrowing status")
	}
	if hasActive {
		return fmt.Errorf("you already have an active borrowing for this book")
	}

	borrowing := &models.Borrowing{
		BookID:  bookID,
		UserID:  userID,
		DueDate: time.Now().AddDate(0, 0, 14),
		Status:  "borrowed",
	}

	return s.borrowingRepo.IssueBook(ctx, borrowing)
}

func (s *borrowingService) ReturnBook(ctx context.Context, bookID, userID int64) error {
	if err := s.borrowingRepo.ReturnBook(ctx, bookID, userID); err != nil {
		return err
	}
	return s.bookRepo.UpdateAvailability(ctx, bookID, 1)
}

func (s *borrowingService) GetMyBorrowings(ctx context.Context, userID int64) ([]models.Borrowing, error) {
	return s.borrowingRepo.GetMyBorrowings(ctx, userID)
}

func (s *borrowingService) GetOverdueBorrowings(ctx context.Context) ([]models.Borrowing, error) {
	return s.borrowingRepo.GetOverdueBorrowings(ctx)
=======
	"database/sql"
)

type BorrowingService struct {
	borrowingRepo *repository.BorrowingRepository
	bookRepo      *repository.BookRepository
}

func NewBorrowingService(db *sql.DB) *BorrowingService {
	return &BorrowingService{
		borrowingRepo: repository.NewBorrowingRepository(db),
		bookRepo:      repository.NewBookRepository(db),
	}
}

func (s *BorrowingService) Borrow(ctx context.Context, userID int64, req models.BorrowRequest) (models.Borrowing, error) {
	if req.BookID == 0 {
		return models.Borrowing{}, errors.New("book_id is required")
	}
	if req.DueDays < 1 || req.DueDays > 60 {
		return models.Borrowing{}, errors.New("due_days must be between 1 and 60")
	}

	book, err := s.bookRepo.GetByID(ctx, req.BookID)
	if err != nil {
		return models.Borrowing{}, errors.New("book not found")
	}
	if book.AvailableCopies <= 0 {
		return models.Borrowing{}, errors.New("book is not available")
	}

	dueDate := time.Now().AddDate(0, 0, req.DueDays)

	borrowing := models.Borrowing{
		BookID:  req.BookID,
		UserID:  userID,          
		DueDate: dueDate,
		Status:  "borrowed",
	}

	created, err := s.borrowingRepo.Create(ctx, borrowing)
	if err != nil {
		return models.Borrowing{}, err
	}

	err = s.bookRepo.UpdateAvailability(ctx, req.BookID, -1)
	if err != nil {
		return models.Borrowing{}, err
	}

	return created, nil
}

func (s *BorrowingService) Return(ctx context.Context, req models.ReturnRequest) error {
	if req.BorrowingID == 0 {
		return errors.New("borrowing_id is required")
	}

	borrowing, err := s.borrowingRepo.GetByID(ctx, req.BorrowingID)
	if err != nil {
		return err
	}

	fine := 0.0
	if borrowing.DueDate.Before(time.Now()) && borrowing.Status == "borrowed" {
		daysOverdue := int(time.Since(borrowing.DueDate).Hours() / 24)
		fine = float64(daysOverdue) * 1.0 
	}

	err = s.borrowingRepo.ReturnBook(ctx, req.BorrowingID, fine)
	if err != nil {
		return err
	}

	return s.bookRepo.UpdateAvailability(ctx, borrowing.BookID, 1)
}

func (s *BorrowingService) GetUserBorrowings(ctx context.Context, userID int64) ([]models.Borrowing, error) {
	return s.borrowingRepo.GetUserBorrowings(ctx, userID)
}

func (s *BorrowingService) GetOverdue(ctx context.Context) ([]models.Borrowing, error) {
	return s.borrowingRepo.GetOverdue(ctx)
}

func (s *BorrowingService) GetByID(ctx context.Context, id int64) (models.Borrowing, error) {
	return s.borrowingRepo.GetByID(ctx, id)
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
}