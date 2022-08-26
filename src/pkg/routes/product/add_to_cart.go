package product

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
	Name string
	Number int
}

func AddToCart(c *fiber.Ctx) error {

	db.Database.Db.SetupJoinTable(&models.User{}, "Product", &models.UserProduct{})
	var product_input Products
	var product models.Product
	var user_products models.UserProduct
	var user models.User
	if err := c.BodyParser(&product_input); err != nil {
		return c.Status(400).JSON(err)
	}
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

	if err := db.Database.Db.First(&product, "name = ? And number >= ?", product_input.Name, product_input.Number).Error; err != nil {
		return c.Status(400).JSON("the product is not available")
	}

	if err:=db.Database.Db.Find(&user_products,"user_id = ? And product_id = ?",id,product.ID).Error; err==nil{
		return c.Status(400).JSON("this product was added to the cart")
	}


	link := models.UserProduct{UserID: int(user.ID), ProductID: int(product.ID), Number:product_input.Number}
	db.Database.Db.Create(&link)

	product.Number-=product_input.Number
	db.Database.Db.Save(&product)
	users, err := GetAllUsers(db.Database.Db)

	return c.Status(200).JSON(users)

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
