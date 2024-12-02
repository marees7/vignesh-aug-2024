package helpers

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Signedvalues struct {
	Email    string
	Name     string
	RoleType string
	RoleId   string
	UserID   int
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateToken(email, name, roletype string, roleid int) (token string, err error) {
	claims := &Signedvalues{
		UserID:   roleid,
		Email:    email,
		Name:     name,
		RoleType: roletype,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(2)).Unix(),
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		fmt.Println("Error occured")
		return
	}
	return token, err
}

func ValidateToken(signedToken string) (claims *Signedvalues, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Signedvalues{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*Signedvalues)
	if !ok {
		msg = "the token is invalid"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}
	return claims, msg
}
