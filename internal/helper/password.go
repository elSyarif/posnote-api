package helper

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

func HashPassword(password string) string {
	salt, _ := strconv.Atoi(os.Getenv("PASS_SALT"))
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
