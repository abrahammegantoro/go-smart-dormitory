package route

import (
	"go-smart-dormitory/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.UserHandlerRead)
	app.Get("/kamar", handler.GetKamarHandlerRead)
	app.Post("/kamar", handler.KamarHandlerCreate)

	app.Put("/kamar/:id", handler.UpdateStatusKamarHandler)
	app.Get("/calon-penghuni", handler.CalonpenghuniHandlerRead)
	app.Post("/penghuni", handler.PenghuniHandlerCreate)
	app.Get("/penghuni", handler.PenghuniHandlerRead)
}
