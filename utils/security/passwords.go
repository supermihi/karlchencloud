package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(password string, hashAndSalt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashAndSalt), []byte(password))
	if err != nil {
		return false
	}
	return true
}
