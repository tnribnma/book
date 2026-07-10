package validators

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Custom validators for the library management system

func RegisterCustomValidators(v *validator.Validate) {
	// ISBN validator (10 or 13 digits)
	v.RegisterValidation("isbn", func(fl validator.FieldLevel) bool {
		isbn := fl.Field().String()
		if isbn == "" {
			return true // optional field
		}
		// Remove hyphens and spaces
		isbn = strings.ReplaceAll(isbn, "-", "")
		isbn = strings.ReplaceAll(isbn, " ", "")

		// Check for 10 or 13 digits
		match, _ := regexp.MatchString(`^\d{10}$|^\d{13}$`, isbn)
		return match
	})

	// Shelf code validator (example: A-12, B-05)
	v.RegisterValidation("shelf", func(fl validator.FieldLevel) bool {
		shelf := fl.Field().String()
		if shelf == "" {
			return true
		}
		match, _ := regexp.MatchString(`^[A-Z]-\d{1,3}$`, shelf)
		return match
	})
}