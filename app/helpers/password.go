package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash ...
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword ...
func VerifyPassword(hashPassowrd, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassowrd), []byte(password))
}
