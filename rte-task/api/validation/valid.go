package validation

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func ValidationFields(user string) (string, error) {
	_, err := mail.ParseAddress(user)
	if err != nil {
		return "", err
	}
	return user, nil
}

func VerifyPassword(first, second string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(second), []byte(first))
	check := true
	msg := ""
	if err != nil {
		msg = "Invalid password"
		check = false
	}
	return check, msg
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
