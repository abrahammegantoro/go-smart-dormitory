package route

import (
	"go-smart-dormitory/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.UserHandlerRead)
	app.Post("/penghuni", handler.PenghuniHandlerCreate)
}
