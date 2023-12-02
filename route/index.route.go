package route

import (
	"go-smart-dormitory/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.UserHandlerRead)

	app.Post("/kamar", handler.KamarHandlerCreate)

	app.Get("/penghuni/:id", handler.PenghuniHandlerReadById)
	app.Delete("/penghuni/:id", handler.DeletePenghuni)
	app.Get("/calon-penghuni", handler.CalonpenghuniHandlerRead)
	app.Get("/penghuni", handler.PenghuniHandlerRead)
}
