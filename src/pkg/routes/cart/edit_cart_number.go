package cart

import (
	"shop/src/pkg/db"
	"shop/src/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func EditCart(c *fiber.Ctx) error {

	var products models.Product
	var user_products models.UserProduct
	var user models.User

	product_id, err := c.ParamsInt("productid")
	number, err := c.ParamsInt("number")

	if err != nil {
		return c.Status(400).JSON(err)
	}

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
		return c.Status(400).JSON("product is not exist")
	}
	if err := db.Database.Db.Find(&user_products, "user_id = ? And product_id = ?", id, products.ID).Error; err != nil {
		return c.Status(400).JSON(user_products)
	}

	db.Database.Db.Model(&user_products).Update("number", number)

	return c.Status(200).JSON("ok")

}
