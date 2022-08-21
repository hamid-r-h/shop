package jwt_handler

import (
	
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(username string) (string, error) {
	
	claims := jwt.MapClaims{}

	claims["exp"] =	jwt.NewNumericDate(time.Unix(time.Now().Add(time.Hour*24).Unix(), 0))
	claims["username"] = username
	new_claim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := new_claim.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
