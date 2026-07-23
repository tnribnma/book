package models

import "time"

type Borrowing struct {
<<<<<<< HEAD
	ID          int64     `json:"id"`
	BookID      int64     `json:"book_id"`
	BookTitle   string    `json:"book_title,omitempty"`
	UserID      int64     `json:"user_id"`
	UserName    string    `json:"user_name,omitempty"`
	IssueDate   time.Time `json:"issue_date"`
	DueDate     time.Time `json:"due_date"`
	ReturnDate  *time.Time `json:"return_date,omitempty"`
	Status      string    `json:"status"`
	FineAmount  float64   `json:"fine_amount"`
	CreatedAt   time.Time `json:"created_at"`
}

type BorrowRequest struct {
	BookID  int64 `json:"book_id" validate:"required"`
	UserID  int64 `json:"user_id" validate:"required"` 
	DueDays int   `json:"due_days" validate:"required,min=1,max=60"`
}
=======
	ID          int64      `json:"id"`
	BookID      int64      `json:"book_id"`
	BookTitle   string     `json:"book_title,omitempty"`
	UserID      int64      `json:"user_id"`
	UserEmail   string     `json:"user_email,omitempty"`
	IssueDate   time.Time  `json:"issue_date"`
	DueDate     time.Time  `json:"due_date"`
	ReturnDate  *time.Time `json:"return_date,omitempty"`
	Status      string     `json:"status"` 
	FineAmount  float64    `json:"fine_amount"`
}

type BorrowRequest struct {
	BookID  int64     `json:"book_id" validate:"required"`
	DueDays int       `json:"due_days" validate:"required,min=1,max=60"` 
}

>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
type ReturnRequest struct {
	BorrowingID int64 `json:"borrowing_id" validate:"required"`
}