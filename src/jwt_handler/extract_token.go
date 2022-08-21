package jwt_handler

import (
	"log"
	"github.com/golang-jwt/jwt/v4"
)

const key = "hamid123456789"

func ExtractTokenData(token string) (string, string) {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, keyFunc)
	log.Println(claims)
	name, ok := claims["username"].(string)
	if !ok {
		return name, "x"
	}
	return name, ""
}

func keyFunc(*jwt.Token) (interface{}, error) {
	return []byte(key), nil
}
