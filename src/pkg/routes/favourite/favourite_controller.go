package favourite

import (
	"shop/src/pkg/db"
	"shop/src/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const key = "secret"

func AddToFavourite(c *fiber.Ctx) error {

	var products models.Product
	var user models.User
	var user_favourites []models.Product
	product_id, err := c.ParamsInt("productid")
	cookie := c.Cookies("jwt")
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cookie, claims, keyFunc)

	claim := token.Claims.(jwt.MapClaims)
	id := claim["id"]
	if err != nil {

	}

	if err := db.Database.Db.First(&user, "ID = ?", id).Error; err != nil {
		return c.Status(400).JSON("first login")
	}

	if err := db.Database.Db.First(&products, "ID = ?", product_id).Error; err != nil {
		return c.Status(400).JSON(product_id)
	}

	user_favourites = append(user_favourites, products)
	user.Favourite = user_favourites
	db.Database.Db.Save(&user)

	return c.Status(200).JSON(user)

}







func keyFunc(*jwt.Token) (interface{}, error) {
	return []byte(key), nil
}
