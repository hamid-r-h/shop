package api

import (
	"shop/src/pkg/routes/auth"
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
	private.Post("/product/addtocart", cart.AddToCart)
	private.Get("/product/getcart", cart.GetCart)
	private.Delete("/product/delete/:id", cart.DeleteCart)
	private.Put("/product/:productid/:number", cart.EditCart)

}
