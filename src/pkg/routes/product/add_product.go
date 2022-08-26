package product

import (
	// "shop/src/db"
	// "shop/src/jwt_handler"
	"shop/src/pkg/db"
	"shop/src/pkg/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const key = "hamid123456789"

func AddProduct(c *fiber.Ctx) error {

	var product models.Product
	var user models.User
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON("sure about input")
	}

	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return c.Status(400).JSON("input in correct format")
	}

	cookie := c.Cookies("jwt")
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cookie, claims, keyFunc)
	claim := token.Claims.(jwt.MapClaims)
	id := claim["id"]

	if err := db.Database.Db.First(&user, "ID = ? And has_access = ? ", id, true).Error; err != nil {
		return c.Status(400).JSON("you have not access")
	}

	if err := db.Database.Db.First(&product, "name = ?", product.Name).Error; err == nil {
		return c.Status(400).JSON(err)
	}

	if product.Category != "mobile" && product.Category != "laptop" {
		return c.Status(400).JSON(product.Category)
	}
	product.UserID = user.ID
	if err := db.Database.Db.Create(&product).Error; err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(product)

}
func keyFunc(*jwt.Token) (interface{}, error) {
	return []byte(key), nil
}
