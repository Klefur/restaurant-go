package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 9)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(hashedPassword string, password string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	check, msg := true, ""
	if err != nil {
		check, msg = false, "invalid password"
	}

	return check, msg
}