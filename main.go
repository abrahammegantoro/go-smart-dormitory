package main

import (
	"go-smart-dormitory/database"
	"go-smart-dormitory/database/migration"
	"go-smart-dormitory/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	migration.RunMigration()

	app := fiber.New()

	route.SetupRoutes(app)

	app.Listen(":8080")
}
