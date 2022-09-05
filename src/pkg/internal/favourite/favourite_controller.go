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
	var user_favourites []models.Product
	product_id, err := c.ParamsInt("productid")
	if err != nil {

	}

	user := c.Locals("user").(models.User)

	if err := db.Database.Db.First(&products, "ID = ?", product_id).Error; err != nil {
		return c.Status(400).JSON(err)
	}

	user_favourites = user.Favourites
	user_favourites = append(user_favourites, products)
	user.Favourites = user_favourites
	db.Database.Db.Save(&user)

	if err := db.Database.Db.Model(&models.User{}).Preload("Favourites").Find(&user, "ID = ? ", user.ID).Error; err != nil {
		return c.Status(400).JSON("ypu have not product")
	}
	return c.Status(200).JSON(user.Favourites)

}

func RemoveFromFavourite(c *fiber.Ctx) error {
	var products models.Product
	product_id, err := c.ParamsInt("productid")
	user := c.Locals("user").(models.User)
	if err != nil {

	}

	if err := db.Database.Db.First(&products, "ID = ?", product_id).Error; err != nil {
		return c.Status(400).JSON(err)
	}

	if err := db.Database.Db.Model(&user).Where("product_id = ? ", product_id).Association("Favourites").Clear(); err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(user)

}

func keyFunc(*jwt.Token) (interface{}, error) {
	return []byte(key), nil
}

