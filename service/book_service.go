package service

import (
	"context"
	"fmt"

	"book-management/models"
	"book-management/repository"
)

type BookService interface {
	CreateBook(ctx context.Context, req models.BookRequest, userID int64) (*models.Book, error)
	GetBook(ctx context.Context, id int64) (*models.Book, error)
	ListBooks(ctx context.Context, filter models.BookFilter) ([]models.Book, error)
	UpdateBook(ctx context.Context, id int64, req models.BookRequest) (*models.Book, error)
	DeleteBook(ctx context.Context, id int64) error
	UpdateAvailability(ctx context.Context, bookID int64, delta int) error
}

type bookService struct {
	bookRepo repository.BookRepository
}

<<<<<<< HEAD
func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (s *bookService) CreateBook(ctx context.Context, req models.BookRequest, userID int64) (*models.Book, error) {
	if req.Quantity <= 0 {
		return nil, fmt.Errorf("quantity must be greater than 0")
	}

	book := &models.Book{
		Title:           req.Title,
		Author:          req.Author,
		ISBN:            req.ISBN,
		CategoryID:      req.CategoryID,
		Publisher:       req.Publisher,
		Edition:         req.Edition,
		PublishedYear:   req.PublishedYear,
		Quantity:        req.Quantity,
		AvailableCopies: req.Quantity,
		Shelf:           req.Shelf,
		Status:          "available",
	}

	if err := s.bookRepo.Create(ctx, book); err != nil {
		return nil, fmt.Errorf("failed to create book: %w", err)
	}

	return book, nil
}

func (s *bookService) GetBook(ctx context.Context, id int64) (*models.Book, error) {
	book, err := s.bookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *bookService) ListBooks(ctx context.Context, filter models.BookFilter) ([]models.Book, error) {
	return s.bookRepo.List(ctx, filter)
}

func (s *bookService) UpdateBook(ctx context.Context, id int64, req models.BookRequest) (*models.Book, error) {
	book, err := s.bookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	book.Title = req.Title
	book.Author = req.Author
	book.ISBN = req.ISBN
	book.CategoryID = req.CategoryID
	book.Publisher = req.Publisher
	book.Edition = req.Edition
	book.PublishedYear = req.PublishedYear
	book.Quantity = req.Quantity
	book.Shelf = req.Shelf

	if err := s.bookRepo.Update(ctx, book); err != nil {
		return nil, fmt.Errorf("failed to update book: %w", err)
	}

	return book, nil
}

func (s *bookService) DeleteBook(ctx context.Context, id int64) error {
	book, err := s.bookRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if book.AvailableCopies < book.Quantity {
		return fmt.Errorf("cannot delete book that is currently borrowed")
	}

	return s.bookRepo.Delete(ctx, id)
}

func (s *bookService) UpdateAvailability(ctx context.Context, bookID int64, delta int) error {
	return s.bookRepo.UpdateAvailability(ctx, bookID, delta)
=======
func (s *BookService) Create(ctx context.Context, req models.BookRequest) (models.Book, error) {
	if req.Quantity < 1 {
		return models.Book{}, errors.New("quantity must be at least 1")
	}

	book := models.Book{
		Title:          req.Title,
		Author:         req.Author,
		ISBN:           req.ISBN,
		CategoryID:     req.CategoryID,
		Publisher:      req.Publisher,
		Edition:        req.Edition,
		PublishedYear:  req.PublishedYear,
		Quantity:       req.Quantity,
		AvailableCopies: req.Quantity,
		Shelf:          req.Shelf,
		Status:         "available",
	}

	return s.repo.Create(ctx, book)
}

func (s *BookService) List(ctx context.Context, filter models.BookFilter, limit, offset int) ([]models.Book, int, error) {
	return s.repo.List(ctx, filter, limit, offset)
}

func (s *BookService) GetByID(ctx context.Context, id int64) (models.Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *BookService) Update(ctx context.Context, id int64, req models.BookRequest) (models.Book, error) {
	return models.Book{}, errors.New("update functionality not fully implemented yet")
}

func (s *BookService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *BookService) Search(ctx context.Context, query string) ([]models.Book, error) {
	return s.repo.Search(ctx, query)
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
}