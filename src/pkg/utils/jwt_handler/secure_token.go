package jwt_handler

import (
	"fmt"
	"shop/src/pkg/db"
	"shop/src/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// SecureAuth returns a middleware which secures all the private routes
func SecureAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("access")
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(accessToken, claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
		fmt.Println(token)

		if err !=nil{
			return c.Status(400).JSON("unathorized")
		}
		if !token.Valid {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					// this is not even a token, we should delete the cookies here
					c.ClearCookie("access", "refresh")
					return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
						"message": "forbidden",
					})
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// Token is either expired or not active yet
					return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
						"message": "forbidden",
					})
				} else {
					// cannot handle this token
					c.ClearCookie("access", "refresh")
					return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
						"message": "forbidden",
					})
				}
			}
		}
		var user models.User
		db.Database.Db.Find(&user,"ID = ?",claims["iss"])
		c.Locals("user",user)
		return c.Next()
	}
}
