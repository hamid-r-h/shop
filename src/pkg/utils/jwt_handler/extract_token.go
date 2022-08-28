package jwt_handler

// import (
// 	"log"

// 	"github.com/golang-jwt/jwt/v4"
// )

// const key = "hamid123456789"

// func ExtractTokenData(cookie interface{}) (interface{}, string) {
// 	claims := jwt.MapClaims{}
// 	token, err := jwt.ParseWithClaims(cookie, claims,keyFunc)
// 	log.Println(err)
// 	if err != nil {
// 		return nil,"error"
// 	}
// 	claim := token.Claims.(jwt.MapClaims)
// 	id := claim["id"]
// 	return id,""
// }

// func keyFunc(*jwt.Token) (interface{}, error) {
// 	return []byte(key), nil
// }
