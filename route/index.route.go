package route

import (
	"go-smart-dormitory/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.UserHandlerRead)
	app.Get("/kamar", handler.GetKamarHandlerRead)

	// kamar
	app.Post("/kamar", handler.KamarHandlerCreate)
	app.Get("/kamar/available", handler.KamarAvailableHandleRead)

	// kontrak
	app.Post("/kontrak", handler.KontrakHandlerCreate)

	// penghuni
	app.Get("/penghuni/:id", handler.PenghuniHandlerReadById)
	app.Delete("/penghuni/:id", handler.PenghuniHandlerDelete)
	app.Get("/calon-penghuni", handler.CalonpenghuniHandlerRead)
	app.Get("/penghuni", handler.PenghuniHandlerRead)
}