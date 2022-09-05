package jwt_handler

import (
	"fmt"
	"shop/src/pkg/db"
	"shop/src/pkg/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(id uint) (string, error) {

	claims := jwt.MapClaims{}

	claims["exp"] = jwt.NewNumericDate(time.Unix(time.Now().Add(time.Second*5).Unix(), 0))
	claims["iss"] = id
	new_claim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	access_token, err := new_claim.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return access_token, nil
}

func GenerateRefreshToken(id uint) (string, error) {

	claims := jwt.MapClaims{}

	claims["exp"] = jwt.NewNumericDate(time.Unix(time.Now().Add(time.Second*24).Unix(), 0))
	claims["iss"] = id
	new_claim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refresh_token, err := new_claim.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return refresh_token, nil
}

func VerifyRefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh")

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(refreshToken, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
	fmt.Println(token)

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
	db.Database.Db.Find(&user, "ID = ?", claims["iss"])
	access_cookie,refresh_cookie,err:=SetCookie(user.ID)
	if(err!=nil){
		c.Status(400).JSON(err)
	}
	c.Cookie(&access_cookie)
	c.Cookie(&refresh_cookie)
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

func JwtGenerateToken(id uint) (string, string, error) {
	access_token, err := GenerateAccessToken(id)
	if err != nil {
		return "", "", err
	}
	refresh_token, err := GenerateRefreshToken(id)
	if err != nil {
		return "", "", err
	}
	return access_token, refresh_token, nil
}
