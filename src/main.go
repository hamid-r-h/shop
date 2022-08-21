package main

import (
	"shop/src/db"
	"shop/src/api"

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
