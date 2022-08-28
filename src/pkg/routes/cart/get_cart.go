package cart

import (
	"log"
	"shop/src/pkg/db"
	"shop/src/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)


func GetCart(c *fiber.Ctx) error {

	var user_products models.UserProduct
	var user models.User
	
	log.Println()
	cookie := c.Cookies("jwt")
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cookie, claims, keyFunc)
	if err != nil {

	}

	claim := token.Claims.(jwt.MapClaims)
	id := claim["id"]

	if err := db.Database.Db.First(&user, "ID = ?", id).Error; err != nil {
		return c.Status(400).JSON("first login")
	}

	if err := db.Database.Db.Find(&user_products, "user_id = ?", id).Error; err != nil {
		return c.Status(400).JSON("ypu have not product")
	}
	if err := db.Database.Db.Model(&models.User{}).Preload("Products").Find(&user,"ID = ? ",id).Error; err != nil {
		return c.Status(400).JSON("ypu have not product")
	}
	GetAllProduct(db.Database.Db)
	return c.Status(200).JSON(user)

}
