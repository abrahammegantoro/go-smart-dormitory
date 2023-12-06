package route

import (
	"go-smart-dormitory/handler"
	"go-smart-dormitory/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// admin
	app.Post("/login", handler.AdminHandlerLogin)

	app.Get("/", middleware.Auth, handler.UserHandlerRead)

	// kamar
	app.Get("/kamar", middleware.Auth, handler.GetKamarHandlerRead)
	app.Post("/kamar", middleware.Auth, handler.KamarHandlerCreate)
	app.Get("/kamar/available", middleware.Auth, handler.KamarAvailableHandleRead)

	// kontrak
	app.Post("/kontrak", middleware.Auth, handler.KontrakHandlerCreate)

	// penghuni
	app.Get("/penghuni/:id", middleware.Auth, handler.PenghuniHandlerReadById)
	app.Delete("/penghuni/:id", middleware.Auth, handler.PenghuniHandlerDelete)
	app.Get("/calon-penghuni", middleware.Auth, handler.CalonpenghuniHandlerRead)
	app.Get("/penghuni", middleware.Auth, handler.PenghuniHandlerRead)
}
