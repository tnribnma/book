package models

type Report struct {
<<<<<<< HEAD
	TotalBooks      int     `json:"total_books"`
	AvailableBooks  int     `json:"available_books"`
	BorrowedBooks   int     `json:"borrowed_books"`
	OverdueBooks    int     `json:"overdue_books"`
	TotalFines      float64 `json:"total_fines"`
	MostBorrowed    string  `json:"most_borrowed,omitempty"`
}

type DashboardStats struct {
	TotalBooks        int `json:"total_books"`
	AvailableBooks    int `json:"available_books"`
	BorrowedBooks     int `json:"borrowed_books"`
	ActiveBorrowings  int `json:"active_borrowings"`
	OverdueBorrowings int `json:"overdue_borrowings"`
	TotalUsers        int `json:"total_users"`
	PendingReservations int `json:"pending_reservations"`
}

type SystemSummary struct {
	TotalBooks        int `json:"total_books"`
	AvailableBooks    int `json:"available_books"`
	ActiveBorrowings  int `json:"active_borrowings"`
	OverdueBorrowings int `json:"overdue_borrowings"`
	TotalUsers        int `json:"total_users"`
=======
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
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
}