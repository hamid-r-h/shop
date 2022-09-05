package auth

import (
	"shop/src/pkg/db"
	"shop/src/pkg/utils/hash"
	"shop/src/pkg/utils/jwt_handler"

	"shop/src/pkg/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username"  valid:"length(4|15),required"`
	Password string `json:"password"  valid:"length(4|15),required"`
}

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON("please sure about input")
	}
	if user.HasAccess != true {
		user.HasAccess = false
	}
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	password, err := hash.HashPassword(user.Password)
	user.Password = password
	if err != nil {
		return c.Status(400).JSON(err)
	}

	if err := db.Database.Db.Create(&user); err.Error != nil {
		return c.Status(200).JSON("already taken")
	}
	access_token, refresh_token, err := jwt_handler.SetCookie(user.ID)
	if err != nil {
		c.Status(400).JSON(err)
	}
	c.Cookie(&access_token)
	c.Cookie(&refresh_token)

	return c.Status(200).JSON(user)

}

func Login(c *fiber.Ctx) error {
	var userinput User
	var user models.User
	if err := c.BodyParser(&userinput); err != nil {
		return c.Status(400).JSON("please sure about input")
	}
	_, err := govalidator.ValidateStruct(userinput)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("please input in  format")
	}

	if err := db.Database.Db.First(&user, "username = ?", userinput.Username).Error; err != nil {
		return c.Status(400).JSON(err)
	}
	check := hash.CheckPasswordHash(userinput.Password, user.Password)
	if !check {
		return c.Status(400).JSON("password incorrect")
	}
	access_token, refresh_token, err := jwt_handler.SetCookie(user.ID)
	if err != nil {
		c.Status(400).JSON(err)
	}
	c.Cookie(&access_token)
	c.Cookie(&refresh_token)
	return c.Status(200).JSON(user)

}
