package validation

import (
	"fmt"
	"regexp"

	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func ValidationSignUp(user models.UserDetails) error {
	if len(user.Name) < 3 {
		return fmt.Errorf(" invalid Name,I need much Longer !Buddy")
	}
	if len(user.Name) > 20 {
		return fmt.Errorf(" invalid Name,I need much shorter, !Buddy")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(user.Email) {
		return fmt.Errorf("invalid Email Id-Your EmailId is not in proper way")
	}

	lowercase := regexp.MustCompile(`[a-z]`)
	uppercase := regexp.MustCompile(`[A-Z]`)
	digit := regexp.MustCompile(`\d`)
	specialChar := regexp.MustCompile(`[@$!%*?&#^()]`)

	if !lowercase.MatchString(user.Password) {
		return fmt.Errorf("invalid Password- Enter Lowercase Letter")
	}
	if !uppercase.MatchString(user.Password) {
		return fmt.Errorf("invalid Password- Enter Uppercase Letter")
	}
	if !digit.MatchString(user.Password) {
		return fmt.Errorf("invalid Password- Enter any Numbers")
	}
	if !specialChar.MatchString(user.Password) {
		return fmt.Errorf("invalid Password- Enter any Special characters")
	}
	if len(user.Password) < 8 {
		return fmt.Errorf("invalid Password- Length should be more than 8")
	}
	if len(user.PhoneNumber) > 10 {
		return fmt.Errorf("invalid phonenumber , Greater than 10 and give properly")
	}
	if len(user.PhoneNumber) < 10 {
		return fmt.Errorf("invalid phonenumber, Less than 10 and Give proeprly")
	}
	if user.RoleType != "USER" && user.RoleType != "ADMIN" {
		return fmt.Errorf("role should be Either USER or ADMIN")
	}
	return nil
}

func VerifyPassword(first, second string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(second), []byte(first))
	check := true
	msg := ""
	if err != nil {
		msg = "Invalid password-Password is Not Match "
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

func ValidationJobPost(post models.JobCreation, paramid int, tokenid int, tokentype string) error {
	if tokentype != "ADMIN" {
		return fmt.Errorf("invalid user-User have not access to create the post")
	}
	if tokenid != paramid {
		return fmt.Errorf("invalid ID,Your Payload ID and RoleId is Mismatching Here,Check It")
	}
	if post.UserID != paramid {
		return fmt.Errorf("invalid ID,Your Payload ID and UserId is Mismatching Here,Check It")
	}
	if len(post.CompanyName) < 3 {
		return fmt.Errorf(" invalid Name,I need much Longer !Buddy")
	}
	if len(post.CompanyName) > 20 {
		return fmt.Errorf(" invalid Name,I need much shorter, !Buddy")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(post.CompanyEmail) {
		return fmt.Errorf("invalid Email Id-Your EmailId is not in proper way")
	}

	if len(post.JobTitle) < 3 {
		return fmt.Errorf(" invalid JobTitle,I need much Longer !Buddy")
	}
	if len(post.JobTitle) > 20 {
		return fmt.Errorf(" invalid JobTitle,I need much shorter, !Buddy")
	}
	if post.JobStatus != "ON GOING" && post.JobStatus != "CLOSED" {
		return fmt.Errorf("invalid jobstatus-It should be either ONGOING or CLOSED")
	}
	if post.JobTime != "PART TIME" && post.JobTime != "FULL TIME" {
		return fmt.Errorf("invalid JobTime-Time should be either PART TIME or FULL TIME")
	}
	if len(post.Description) < 3 {
		return fmt.Errorf(" invalid Description,I need much Longer !Buddy")
	}
	if len(post.Skills) < 3 {
		return fmt.Errorf(" invalid Skills,I need much Longer !Buddy")
	}
	if post.Vacancy < 0 {
		return fmt.Errorf(" invalid Vacnacy,You can Enter some Vacancy Here")
	}
	if len(post.Country) < 3 {
		return fmt.Errorf(" invalid Country,I need much Longer !Buddy")
	}
	if len(post.Address.Street) < 3 {
		return fmt.Errorf("invalid Street,I need much Longer !Buddy")
	}
	if len(post.Address.City) < 3 {
		return fmt.Errorf("invalid city,I need much Longer !Buddy")
	}
	if len(post.Address.State) < 3 {
		return fmt.Errorf("invalid city,I need much Longer !Buddy")
	}

	if len(post.Address.ZipCode) < 6 {
		return fmt.Errorf("invalid pincode,I need much Longer,!Buddy")
	}
	if len(post.Address.ZipCode) > 6 {
		return fmt.Errorf("invalid pincode,I need only Six Numbers,!Buddy")
	}
	return nil
}

func ValidationUserJob(user models.UserJobDetails, tokentype string, tokenid int, parmid int) error {
	if tokentype != "USER" {
		return fmt.Errorf("invalid Admin-Admin cannot have access to apply the post")
	}
	if tokenid != parmid {
		return fmt.Errorf("invalid ID,Your Payload ID and RoleId is Mismatching Here,Check It")
	}
	if user.UserID != parmid {
		return fmt.Errorf("invalid ID,Your Payload ID and UserId is Mismatching Here,Check It")
	}

	if user.Experience < 0 {
		return fmt.Errorf("invalid experience,Experience shoud be in Two Digits")
	}
	if len(user.Skills) < 3 {
		return fmt.Errorf(" invalid Skills,I need much Longer !Buddy")
	}
	if len(user.Language) < 3 {
		return fmt.Errorf(" invalid Skills,I need much Longer !Buddy")
	}
	if len(user.Country) < 1 {
		return fmt.Errorf(" invalid Country name,I need much Longer !Buddy")
	}
	if len(user.Country) > 20 {
		return fmt.Errorf(" invalid Country,I need much shorter, !Buddy")
	}
	return nil
}

func ValidationAdminFields(post models.JobCreation, tokentype string, tokenid int, parmid int) error {
	if tokentype != "ADMIN" {
		return fmt.Errorf("invalid user-User have not access to view this details")
	}
	if tokenid != parmid {
		return fmt.Errorf("invalid ID,Your Payload ID and RoleId is Mismatching Here,Check It")
	}
	if post.UserID != parmid {
		return fmt.Errorf("invalid ID,Your Payload ID and UserId is Mismatching Here,Check It")
	}
	return nil
}

func ValidationCheck(tokentype string, tokenid int, parmid int) error {
	if tokentype != "ADMIN" {
		return fmt.Errorf("invalid user-User have not access to view this details")
	}
	if tokenid != parmid {
		return fmt.Errorf("invalid ID,Your Payload ID and RoleId is Mismatching Here,Check It")
	}
	return nil
}

// func ValidateRoletype(usertype string) error {
// 	if usertype != "ADMIN" {
// 		return fmt.Errorf("invalid Authorization,Only admin and users can view this Posts")
// 	} else if usertype != "USER" {
// 		return fmt.Errorf("invalid Authorization,Only admin and users can view this Posts")
// 		// if usertype != "USER" {
// 	} else {
// 		return nil
// 	}
// }
