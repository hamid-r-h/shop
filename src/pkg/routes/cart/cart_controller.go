package cart

import (
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
	if err := db.Database.Db.Model(&models.User{}).Preload("Products").Find(&user, "ID = ? ", id).Error; err != nil {
		return c.Status(400).JSON("ypu have not product")
	}
	GetAllProduct(db.Database.Db)
	return c.Status(200).JSON(user.Products)

}
