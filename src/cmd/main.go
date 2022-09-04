package main

import (
	"shop/src/pkg/api"
	"shop/src/pkg/db"


	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.ConnectDb()
	app := fiber.New()

	api.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
