package pagination

import (
	"fmt"
	"shop/src/pkg/db"
	"shop/src/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GeneratePagInation(c *fiber.Ctx) models.PagInation {

	var pagination_model models.PagInation
	current_page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	pagination_model.Limit = limit
	pagination_model.Page = current_page
	pagination_model.Sort = c.Query("sort")
	return pagination_model
}

func GetProduct(pagination *models.PagInation, category string, subcategory string) (*[]models.Product, string) {
	var products []models.Product
	if pagination.Limit > 20 || pagination.Limit < 2 {
		return nil, "limit is out of range"
	}
	offset := pagination.Limit * (pagination.Page - 1)
	queryBuider := db.Database.Db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	if category != "all" {
		if subcategory != "all" {
			queryBuider.Model(&models.Product{}).Where("category = ? And sub_category = ?", category, subcategory).Find(&products)
			fmt.Println(subcategory)
		} else {
			queryBuider.Model(&models.Product{}).Where("category = ?", category).Find(&products)
		}
	} else {
		queryBuider.Model(&models.Product{}).Find(&products)
	}
	return &products, ""
}

func GetComment(pagination *models.PagInation) (*[]models.Comment, string) {
	var comments []models.Comment
	if pagination.Limit > 20 || pagination.Limit < 2 {
		return nil, "limit is out of range"
	}
	offset := pagination.Limit * (pagination.Page - 1)
	queryBuider := db.Database.Db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	queryBuider.Model(&models.Comment{}).Find(&comments)
	return &comments, ""
}