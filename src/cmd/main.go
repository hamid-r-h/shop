package main

import (
	"shop/src/pkg/api"
	"shop/src/pkg/db"


	// routes "shop/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.ConnectDb()
	app := fiber.New()

	api.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
