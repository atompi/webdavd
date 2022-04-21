package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordHasher(inputStr string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(inputStr), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func PasswordVerifier(hashedStr string, inputStr string) (result bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(inputStr))
	if err != nil {
		result = false
		return
	}
	result = true
	return
}
