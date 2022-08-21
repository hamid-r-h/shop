package product

import (
	// "shop/src/db"
	// "shop/src/jwt_handler"
	"shop/src/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)
const key = "hamid123456789"

func AddProduct(c *fiber.Ctx) error {

	var product models.Product
	//  var user models.User
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON("sure about input")
	}

	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return c.Status(400).JSON("input in correct format")
	}

	token := c.Cookies("jwt")
	claims := jwt.MapClaims{}
	tk,err:=jwt.ParseWithClaims(token, claims, keyFunc)
	claim :=tk.Claims.(jwt.MapClaims)
	name := claim["username"]
	return c.Status(200).JSON(name)
	
	// if err := db.Database.Db.First(&user, "username = ? And access = ? ", name, true).Error; err != nil {
	// 	return c.Status(400).JSON(err)
	// }

	// if err := db.Database.Db.First(&product, "name = ?", product.Name).Error; err == nil {
	// 	return c.Status(400).JSON(err)
	// }

	// if product.Category != "mobile" || product.Category != "laptop" {
	// 	return c.Status(400).JSON("category invalid")
	// }
	// db.Database.Db.Create(&product)

	// return c.Status(200).JSON("succesfully addProduct")

}
func keyFunc(*jwt.Token) (interface{}, error) {
	return []byte(key), nil
}
