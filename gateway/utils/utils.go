package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TryLogin(Username string, Password string) string{
	hash, _ := HashPassword(password)
	// check database
	
}

func TryRegister(Username string, Password string) string{
	hash, _ := HashPassword(password)

}