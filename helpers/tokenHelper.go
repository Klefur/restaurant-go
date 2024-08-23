package helpers

import (
	"go-restaurant/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type TokenClaims struct {
	ID       uint
	Username string
	Email    string
	jwt.StandardClaims
}

func GenerateToken(user models.User) (token string, err error) {
	godotenv.Load()
	SECRET_KEY := os.Getenv("SECRET_KEY")

	claims := TokenClaims{
		user.ID,
		user.First_name + " " + user.Last_name,
		user.Email,
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

func ValidateToken(token string) (claims *TokenClaims, msg string) {
	godotenv.Load()
	SECRET_KEY := os.Getenv("SECRET_KEY")

	tkn, err := jwt.ParseWithClaims(
		token,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	claims, ok := tkn.Claims.(*TokenClaims)

	if !ok {
		return claims, "Invalid token"
	}

	if err != nil {
		return claims, err.Error()
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return claims, "Token has expired"
	}

	return claims, ""

}
