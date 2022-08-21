package auth

import (
    "time"
	"shop/src/db"
	"shop/src/jwt_handler"
	"shop/src/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username  string `json:"user_name"  valid:"length(4|15),required"`
	Password  string `json:"user_pass"  valid:"length(4|15),required"`
	HasAccess bool   `json:"Access"`
	Email     string `json:"user_email"  valid:"email,required"`
}

type User struct {
	Username string `json:"user_name"  valid:"length(4|15),required"`
	Password string `json:"user_pass"  valid:"length(4|15),required"`
}

func Register(c *fiber.Ctx) error {
	var user Users
	var users models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON("please sure about input")
	}
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return c.Status(400).JSON("please input in correct format")
	}

	if err := db.Database.Db.First(&user, "username = ?", user.Username).Error; err == nil {
		return c.Status(400).JSON(err)
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	token, err := jwt_handler.GenerateAccessToken(user.Username)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	users.Password = password
	users.Email = user.Email
	users.Username = user.Username
	users.HasAccess = user.HasAccess

	db.Database.Db.Create(&users)
	c.Set("Authorization", token)
	c.Set("Content-Type", "application/json")

	return c.Status(200).JSON("succesfully signup")

}

func Login(c *fiber.Ctx) error {
	var userinput User
	var user models.User
	if err := c.BodyParser(&userinput); err != nil {
		return c.Status(400).JSON("please sure about input")
	}
	_, err := govalidator.ValidateStruct(userinput)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("please input in correct format")
	}

	if err := db.Database.Db.First(&user, "username = ?", userinput.Username).Error; err != nil {
		return c.Status(400).JSON(err)
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(userinput.Password)); err != nil {
		return c.Status(400).JSON(err)
	}

	token, err := jwt_handler.GenerateAccessToken(user.Username)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.Status(200).JSON("successfully login")

}
