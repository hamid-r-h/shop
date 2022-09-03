package api

import (
	"shop/src/pkg/routes/auth"
	"shop/src/pkg/routes/favourite"
	"shop/src/pkg/routes/cart"
	"shop/src/pkg/routes/product"
	"shop/src/pkg/utils/jwt_handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	routes := app.Group("/api")
	routes.Post("/user/register", auth.Register)
	routes.Post("/user/login", auth.Login)
	private := routes.Group("/private")
	private.Use(jwt_handler.SecureAuth())
	private.Post("/product/addproduct", product.AddProduct)
	private.Post("/cart/add", cart.AddToCart)
	private.Get("/cart/get", cart.GetCart)
	private.Delete("/cart/delete/:id", cart.DeleteCart)
	private.Put("/cart/edit/:productid/:number", cart.EditCart)
	private.Put("/product/addfavourite/:productid",favourite.AddToFavourite)
}
