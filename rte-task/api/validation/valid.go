package validation

import (
	"fmt"
	"regexp"

	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func ValidationSignUp(user models.UserDetails) error {
	if len(user.Name) == 0 {
		return fmt.Errorf(" missing Name,I need much Longer !Buddy")
	}
	if len(user.Name) > 30 {
		return fmt.Errorf(" your Name should be larger,I need much shorter, !Buddy")
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

	if post.DomainID != paramid {
		return fmt.Errorf("invalid ID,Your Payload ID and UserId is Mismatching Here,Check It")
	}

	if len(post.CompanyName) == 0 {
		return fmt.Errorf("missing CompanyName,I need much Longer !Buddy")
	}

	if len(post.CompanyName) > 30 {
		return fmt.Errorf("your CompanyName should be larger,I need much shorter, !Buddy")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(post.CompanyEmail) {
		return fmt.Errorf("invalid Email Id-Your EmailId is not in proper way")
	}

	if len(post.JobTitle) == 0 {
		return fmt.Errorf(" missing JobTitle,I need much Longer !Buddy")
	}

	if len(post.JobTitle) > 20 {
		return fmt.Errorf(" your JobTitle should be larger,I need much shorter, !Buddy")
	}
	if post.JobStatus != "IN PROGRESS" && post.JobStatus != "COMPLETED" && post.JobStatus != "ON GOING" {
		return fmt.Errorf("invalid jobstatus-It should be either IN PROGRESS OR ON GOING")
	}
	if post.JobTime != "PART TIME" && post.JobTime != "FULL TIME" {
		return fmt.Errorf("invalid JobTime-Time should be either PART TIME or FULL TIME")
	}
	if len(post.Description) == 0 {
		return fmt.Errorf(" missing Description,I need much Longer !Buddy")
	}
	if len(post.Description) < 3 {
		return fmt.Errorf(" your Description is very small,I need much Longer !Buddy")
	}
	if len(post.Skills) == 0 {
		return fmt.Errorf(" missing Skills,I need much Longer !Buddy")
	}
	if post.Vacancy < 0 {
		return fmt.Errorf(" missing Vacnacy,You can Enter some Vacancy Here")
	}
	if len(post.Country) == 0 {
		return fmt.Errorf(" missing Country,I need much Longer !Buddy")
	}
	if len(post.Address.Street) == 0 {
		return fmt.Errorf("missing Street,I need much Longer !Buddy")
	}
	if len(post.Address.City) == 0 {
		return fmt.Errorf("missing city,I need much Longer !Buddy")
	}
	if len(post.Address.State) == 0 {
		return fmt.Errorf("missing state,I need much Longer !Buddy")
	}

	if len(post.Address.ZipCode) == 0 {
		return fmt.Errorf("missing ZipCode,I need much Longer,!Buddy")
	}
	if len(post.Address.ZipCode) < 6 {
		return fmt.Errorf("your ZipCode is smaller,I need much Longer,!Buddy")
	}
	if len(post.Address.ZipCode) > 6 {
		return fmt.Errorf("your ZipCode is larger,I need only Six Numbers,!Buddy")
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

	if user.UserId != parmid {
		return fmt.Errorf("invalid ID,Your Payload ID and UserId is Mismatching Here,Check It")
	}

	if user.Experience < 0 {
		return fmt.Errorf("missing experience,Experience shoud be in Two Digits")
	}

	if len(user.Skills) == 0 {
		return fmt.Errorf(" missing Skills,I need your skills !Buddy")
	}

	if len(user.Language) == 0 {
		return fmt.Errorf(" missing Language,I need your Languages !Buddy")
	}

	if len(user.Country) == 0 {
		return fmt.Errorf(" missing Country ,I need your country !Buddy")
	}
	if len(user.JobRole) == 0 {
		return fmt.Errorf(" missing JobRole ,I need some Jobrole here !Buddy")
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
	if post.DomainID != parmid {
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

func ValidationUpdatePost(post models.JobCreation) error {
	if post.JobStatus != "COMPLETED" && post.JobStatus != "ON GOING" {
		return fmt.Errorf("invalid jobstatus,Only Completed or On Going only")
	}
	if post.Vacancy != 0 {
		return fmt.Errorf("invalid vacancy,check it properly")
	}
	return nil
}
