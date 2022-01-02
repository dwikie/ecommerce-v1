package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 13)

	return string(passwordHash)
}

func ComparePassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
