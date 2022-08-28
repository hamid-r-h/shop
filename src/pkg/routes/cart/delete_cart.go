package cart

import (
	"shop/src/pkg/db"
	"shop/src/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func DeleteCart(c *fiber.Ctx) error {

	var products models.Product
	var user_products models.UserProduct
	var user models.User

	product_id, err := c.ParamsInt("id")
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
	if err := db.Database.Db.First(&products, "ID = ?", product_id).Error; err != nil {
		return c.Status(400).JSON("first login")
	}
	if err := db.Database.Db.Where("user_id = ? And product_id = ?", id, products.ID).Delete(&user_products); err != nil {
		return c.Status(400).JSON(err)

	}

	return c.Status(200).JSON("ok")

}
