package apputil

import (
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 14

// HashPassword hashes a password.
func HashPassword(password string) (string, error) {
	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassowrd), nil
}

// ComparePassword compares a password with a hash.
func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
