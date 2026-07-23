package models

<<<<<<< HEAD
import "time"

type Category struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
=======
type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description,omitempty" validate:"omitempty,max=500"`
>>>>>>> 9643364dd4f1350f52d70f9a28ef341da82933d8
}