package handler

import (
	"github.com/gofiber/fiber/v2"
)

func UserHandlerRead(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World 👋!")
}
