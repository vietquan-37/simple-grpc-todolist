package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUserName    = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName    = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidPhoneNumber = regexp.MustCompile(`^\d{10,11}$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain between from %d-%d characters", minLength, maxLength)
	}
	return nil
}
func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidUserName(value) {
		return fmt.Errorf("must contain only lowercase letters, digit or underscore")
	}
	return nil
}
func ValidatePhoneNumber(value string) error {
	if !isValidPhoneNumber(value) {
		return fmt.Errorf("must be contain all digit from 10 to 11")
	}
	return nil
}
func ValidatePassword(password string) error {
	return ValidateString(password, 6, 10)
}
func ValidateEmail(email string) error {
	if err := ValidateString(email, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}
func ValidateFullname(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters and space")
	}
	return nil
}
