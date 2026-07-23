package models

import "time"

type Book struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	ISBN            string    `json:"isbn,omitempty"`
	CategoryID      *int64    `json:"category_id,omitempty"`
	CategoryName    string    `json:"category_name,omitempty"`
	Publisher       string    `json:"publisher,omitempty"`
	Edition         string    `json:"edition,omitempty"`
	PublishedYear   int       `json:"published_year,omitempty"`
	Quantity        int       `json:"quantity"`
	AvailableCopies int       `json:"available_copies"`
	Shelf           string    `json:"shelf,omitempty"`
<<<<<<< HEAD
	Status          string    `json:"status"`
=======
	Status          string    `json:"status"` 
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
	CreatedAt       time.Time `json:"created_at"`
}

type BookRequest struct {
	Title         string  `json:"title" validate:"required,min=1,max=255"`
	Author        string  `json:"author" validate:"required,min=1,max=255"`
<<<<<<< HEAD
	ISBN          string  `json:"isbn,omitempty" validate:"omitempty,isbn"`
=======
	ISBN          string  `json:"isbn,omitempty" validate:"omitempty,len=10|len=13"`
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
	CategoryID    *int64  `json:"category_id,omitempty"`
	Publisher     string  `json:"publisher,omitempty" validate:"omitempty,max=100"`
	Edition       string  `json:"edition,omitempty" validate:"omitempty,max=50"`
	PublishedYear int     `json:"published_year,omitempty" validate:"omitempty,min=1000,max=2100"`
	Quantity      int     `json:"quantity" validate:"required,min=1"`
<<<<<<< HEAD
	Shelf         string  `json:"shelf,omitempty" validate:"omitempty,shelf"`
=======
	Shelf         string  `json:"shelf,omitempty" validate:"omitempty,max=50"`
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
}

type BookFilter struct {
	Search   string `json:"search,omitempty"`
	Category int64  `json:"category,omitempty"`
	Author   string `json:"author,omitempty"`
	Status   string `json:"status,omitempty"`
}