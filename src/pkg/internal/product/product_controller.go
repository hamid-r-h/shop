package product

import (
	// "encoding/json"
	"shop/src/pkg/db"
	"shop/src/pkg/internal/pagination"
	"shop/src/pkg/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const key = "hamid123456789"

func AddProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON("sure about input")
	}
	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return c.Status(400).JSON("input in correct format")
	}

	user := c.Locals("user").(models.User)

	if !user.HasAccess {
		return c.Status(400).JSON(err)
	}

	if err := db.Database.Db.First(&product, "name = ?", product.Name).Error; err == nil {
		return c.Status(400).JSON(err)
	}

	if product.Category != "mobile" && product.Category != "laptop" {
		return c.Status(400).JSON(product.Category)
	}
	if err := db.Database.Db.Create(&product).Error; err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(product)
}
func keyFunc(*jwt.Token) (interface{}, error) {
	return []byte(key), nil
}

func GetAllProduct(c *fiber.Ctx) error {
	pagination_model := pagination.GeneratePagInation(c)
	category:="all"
	subcategory:="all"
	products, ok := pagination.GetProduct(&pagination_model,category,subcategory)
	if ok != "" {
		return c.Status(fiber.StatusBadRequest).JSON(ok)
	}
	return c.Status(400).JSON(products)

}

func GetByCategory(c *fiber.Ctx) error {
	category := c.Params("category")
	pagination_model := pagination.GeneratePagInation(c)
	subcategory := "all"
	products, ok := pagination.GetProduct(&pagination_model, category, subcategory)
	if ok != "" {
		return c.Status(fiber.StatusBadRequest).JSON(ok)
	}
	return c.Status(200).JSON(products)
}

func GetBySubCategory(c *fiber.Ctx) error {
	category := c.Params("category")
	subcategory := c.Params("subcategory")
	pagination_model := pagination.GeneratePagInation(c)
	products, ok := pagination.GetProduct(&pagination_model, category, subcategory)
	if ok != "" {
		return c.Status(fiber.StatusBadRequest).JSON(ok)
	}
	return c.Status(200).JSON(products)
}
func GetByName(c *fiber.Ctx) error {
	var product models.Product
	var product_name ProductName
	if err := c.BodyParser(&product_name.Name); err != nil {
		c.Status(400).JSON(err)
	}
	if err := db.Database.Db.Find(&product, "name = ?", product_name.Name).Error; err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(product)
}

type ProductName struct {
	Name string
}
