package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// >> returns hashed password from string password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),
		bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("@: Error: hasing failed: %w", err)
	}
	return string(hashedPassword), nil
}

// >>  compares hashed, and non-hashed passwords
func CheckPassword(password string, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed),
		[]byte(password))
}
