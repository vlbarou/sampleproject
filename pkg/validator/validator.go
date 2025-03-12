package validator

import (
	"regexp"
)

// ValidateEmail checks if the given email has a valid format
func ValidateEmail(email string) bool {
	// Regular expression for validating email
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}

// ValidateUsername checks if the username contains only uppercase/lowercase letters and numbers.
func ValidateUsername(username string) bool {
	// Regular expression: only A-Z, a-z, and 0-9 allowed
	var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return usernameRegex.MatchString(username)
}
