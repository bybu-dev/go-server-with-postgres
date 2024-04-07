package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10);
	if (err != nil) {
		return "", fmt.Errorf("could not hash password %w", err);
	}

	return string(hashPassword), nil;
}

func CompareHashPassword(password string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}