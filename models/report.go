package models

type Report struct {
	TotalBooks       int     `json:"total_books"`
	AvailableBooks   int     `json:"available_books"`
	BorrowedBooks    int     `json:"borrowed_books"`
	OverdueBooks     int     `json:"overdue_books"`
	TotalFines       float64 `json:"total_fines"`
	MostBorrowedBook string  `json:"most_borrowed_book,omitempty"`
}

type DashboardStats struct {
	TotalUsers    int `json:"total_users"`
	ActiveUsers   int `json:"active_users"`
	TotalBorrowings int `json:"total_borrowings"`
}