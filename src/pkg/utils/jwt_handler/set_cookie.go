package jwt_handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)


func SetCookie(id uint) (fiber.Cookie,fiber.Cookie,error) {

	access_token, refresh_token, err := JwtGenerateToken(id)
	if err != nil {
		return fiber.Cookie{},fiber.Cookie{},err
	}

	access_cookie := fiber.Cookie{
		Name:     "access",
		Value:    access_token,
		Expires:  time.Now().Add(time.Minute * 5),
		HTTPOnly: false,
	}

	refresh_cookie := fiber.Cookie{
		Name:     "refresh",
		Value:    refresh_token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: false,
	}
	return access_cookie,refresh_cookie,nil


}