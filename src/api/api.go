package api

import (
	"shop/src/routes/auth"
	"shop/src/routes/product"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    routes:=app.Group("/api")
	routes.Post("/user/register", auth.Register)
	routes.Post("/user/login", auth.Login)
	routes.Post("/product/addproduct",product.AddProduct)
}
