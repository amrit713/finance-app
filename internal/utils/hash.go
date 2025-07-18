package utils

import (
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), bcrypt.DefaultCost)

	return string(bytes), err

}

func ComparePasswordHash(password string, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(strings.TrimSpace(password)))
	if err != nil {
		log.Println("error", err)
	}
	return err == nil
}
