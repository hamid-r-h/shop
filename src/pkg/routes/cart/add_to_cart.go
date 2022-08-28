package cart

import (
	// "shop/src/db"
	// "shop/src/jwt_handler"
	"log"
	"shop/src/pkg/db"
	"shop/src/pkg/models"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Products struct {
	Name   string
	Number int
}

const key = "secret"

func AddToCart(c *fiber.Ctx) error {

	var get_detail Products
	var products models.Product
	var user models.User
	var user_products models.UserProduct
	if err := c.BodyParser(&get_detail); err != nil {
		return c.Status(400).JSON(err)
	}
	log.Println()
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

	if err := db.Database.Db.First(&products, "name = ? And number >= ?", get_detail.Name, get_detail.Number).Error; err != nil {
		return c.Status(400).JSON(err)
	}

	products.Number -= get_detail.Number
	db.Database.Db.Save(&products)
	tmp_product := user.Products
	tmp_product = append(tmp_product, products)
	user.Products = tmp_product
	db.Database.Db.Save(&user)
	if err := db.Database.Db.Find(&user_products, "user_id = ? And product_id = ?", id, products.ID).Error; err != nil {
		return c.Status(400).JSON(err)
	}
	user_products.Number = get_detail.Number
	db.Database.Db.Save(&user_products)
	GetAllUsers(db.Database.Db)

	return c.Status(200).JSON(user_products)

}
func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	err := db.Model(&models.User{}).Preload("Products").Find(&users).Error
	return users, err
}
func GetAllProduct(db *gorm.DB) ([]models.Product, error) {
	var products []models.Product
	err := db.Model(&models.Product{}).Preload("Users").Find(&products).Error
	return products, err
}
func keyFunc(*jwt.Token) (interface{}, error) {
	return []byte(key), nil
}
