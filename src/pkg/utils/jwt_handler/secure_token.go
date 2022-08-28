package jwt_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// SecureAuth returns a middleware which secures all the private routes
func SecureAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("jwt")
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(accessToken, claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})

		if !token.Valid {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					// this is not even a token, we should delete the cookies here
					c.ClearCookie("jwt", "refresh_token")
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
					c.ClearCookie("access_token", "refresh_token")
					return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
						"message": "forbidden",
					})
				}
			}
		}

		c.Locals("id", claims["id"])
		return c.Next()
	}
}
