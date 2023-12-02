package main

import (
	"go-smart-dormitory/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	route.SetupRoutes(app)

	app.Listen(":8080")
}
