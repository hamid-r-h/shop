package api

import (
	"shop/src/pkg/internal/auth"
	"shop/src/pkg/internal/cart"
	"shop/src/pkg/internal/favourite"
	"shop/src/pkg/internal/product"
	"shop/src/pkg/utils/jwt_handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	internal := app.Group("/api")
	internal.Post("/user/register", auth.Register)
	internal.Post("/user/login", auth.Login)
	internal.Get("/product/getall", product.GetAllProduct)
	internal.Get("/product/:category", product.GetByCategory)
	internal.Get("/product/:category/:subcategory",product.GetBySubCategory)

	private := internal.Group("/private")
	private.Use(jwt_handler.SecureAuth())
	private_cart := private.Group("/cart")
	private_cart.Post("/add", cart.AddToCart)
	private_cart.Get("/get", cart.GetCart)
	private_cart.Delete("/delete/:id", cart.DeleteFromCart)
	private_cart.Put("/edit/:productid/:number", cart.EditCart)
	private_product := private.Group("product")
	private_product.Post("/addproduct", product.AddProduct)
	private_product.Put("/addfavourite/:productid", favourite.AddToFavourite)
	private_product.Delete("/removefavourite/:productid", favourite.RemoveFromFavourite)
	
}
