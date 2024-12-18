package validation

import (
	"fmt"
	"regexp"

	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// signup their each fields
func ValidationSignUp(user models.UserDetails) error {
	if len(user.Name) == 0 {
		return fmt.Errorf(" missing Name,I need much Longer !Buddy")
	}

	if len(user.Name) > 30 {
		return fmt.Errorf(" your Name will be Larger,I need much shorter, !Buddy")
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
		return fmt.Errorf("invalid Password- Enter any Lowercase Letter")
	}
	if !uppercase.MatchString(user.Password) {
		return fmt.Errorf("invalid Password- Enter any Uppercase Letter")
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

// verify their password is match with signup password
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

// Hashing the password here
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Valid their Job Post with Fields
func ValidationJobPost(post models.JobCreation, roleID int, roleType string) error {
	if roleType != "ADMIN" {
		return fmt.Errorf("invalid user-User have not access to create the post")
	}

	if post.AdminID != roleID {
		return fmt.Errorf("invalid ID,Your Role ID and User ID is Mismatching Here,Check It")
	}

	if len(post.CompanyName) == 0 {
		return fmt.Errorf("missing CompanyName,I need your CompanyName Here !Buddy")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(post.CompanyEmail) {
		return fmt.Errorf("invalid Email Id-Your EmailId is not in proper way")
	}

	if len(post.JobTitle) == 0 {
		return fmt.Errorf(" missing JobTitle,I need much Longer !Buddy")
	}

	if len(post.JobTitle) > 20 {
		return fmt.Errorf(" your JobTitle will be larger,I need much shorter, !Buddy")
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

	if len(post.Experience) == 0 {
		return fmt.Errorf("missing Experience,It have some experience details")
	}

	if len(post.Experience) < 3 {
		return fmt.Errorf("your experience should be much greater")
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
		return fmt.Errorf("your ZipCode is Smaller,I need only Six Digits,!Buddy")
	}

	if len(post.Address.ZipCode) > 6 {
		return fmt.Errorf("your ZipCode is  Larger,I need only Six Digits,!Buddy")
	}
	return nil
}

// Valid their User JobPost with Fields
func ValidationUserJob(user models.UserJobDetails, roleType string, userID int) error {
	if roleType != "USER" {
		return fmt.Errorf("invalid Admin-Admin cannot have access to apply the post")
	}

	if user.UserId != userID {
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

// valid their JobFields in JobPosts
func ValidationUpdatePost(post models.JobCreation, roleType string, roleID int) error {
	if roleType != "ADMIN" {
		return fmt.Errorf("invalid user-User have not access to view this details")
	}

	if post.AdminID != roleID {
		return fmt.Errorf("you are not authorized to update this job post")
	}
	if post.JobStatus != "COMPLETED" && post.JobStatus != "ON GOING" {
		return fmt.Errorf("invalid jobstatus,Only Completed or On Going only")
	}
	if post.Vacancy != 0 {
		return fmt.Errorf("invalid vacancy,check it properly")
	}
	return nil
}

// check their roles by users
func ValidateUserType(roleType string) error {
	if roleType != "USER" {
		return fmt.Errorf("invalid admin-Admin have not to access this details")
	}
	return nil
}

// check their roles by admin or users
func ValidateRoleType(roleType string) error {
	if roleType != "ADMIN" {
		return fmt.Errorf("invalid user-User have not access to view this details")
	}
	return nil
}
