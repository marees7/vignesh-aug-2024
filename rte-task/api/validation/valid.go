package validation

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func ValidationFields(user string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(user)
}

func ValidationName(name string) error {
	if len(name) < 3 {
		return fmt.Errorf(" need name much Longer, !Buddy")
	}
	if len(name) > 20 {
		return fmt.Errorf(" need name much shorter, !Buddy")
	}
	return nil
}
func ValidationPassword(passwordRegex string) error {
	lowercase := regexp.MustCompile(`[a-z]`)
	uppercase := regexp.MustCompile(`[A-Z]`)
	digit := regexp.MustCompile(`\d`)
	specialChar := regexp.MustCompile(`[@$!%*?&#^()]`)

	if !lowercase.MatchString(passwordRegex) {
		return fmt.Errorf("enter any Lowercase Letter")
	}
	if !uppercase.MatchString(passwordRegex) {
		return fmt.Errorf("enter any Uppercase Letter")
	}
	if !digit.MatchString(passwordRegex) {
		return fmt.Errorf("enter any Numbers")
	}
	if !specialChar.MatchString(passwordRegex) {
		return fmt.Errorf("enter any Special characters")
	}
	if len(passwordRegex) < 8 {
		return fmt.Errorf("password length should be more than 8")
	}
	return nil
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

func ValidationPhoneNumber(phonenumber string) error {
	if len(phonenumber) > 10 {
		return fmt.Errorf("phonenumber is Greater than 10 and give properly")
	}
	if len(phonenumber) < 10 {
		return fmt.Errorf("phoneNumber is Less than 10 and Give proeprly")
	}
	return nil
}
