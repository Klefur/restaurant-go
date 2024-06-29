package helpers

import (
	"go-restaurant/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)	

type TokenClaims struct {
	ID uint
	Username string
	Email string
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateToken(user models.User) (token string, err error) {

	claims := TokenClaims{
		user.ID,
		*user.First_name + " " + *user.Last_name,
		*user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string) (isValid bool, msg string) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	claims, ok := tkn.Claims.(*TokenClaims)

	if !ok || !tkn.Valid {
		return false, "Invalid token"
	}

	if err != nil {
		return false, err.Error()
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return false, "Token has expired"
	}

	return true, ""

}