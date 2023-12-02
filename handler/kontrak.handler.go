package handler

import (
	"go-smart-dormitory/database"
	"go-smart-dormitory/model/entity"
	"log"
	"github.com/gofiber/fiber/v2"
)

func KontrakHandlerCreate(ctx *fiber.Ctx) error {
	var kontrak entity.Kontrak

	if err := ctx.BodyParser(&kontrak); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	result := database.DB.Create(&kontrak)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	// update status kamar
	var kamar entity.Kamar
	database.DB.First(&kamar, kontrak.KamarID)
	kamar.Status = "Booked"
	database.DB.Save(&kamar)

	return ctx.JSON(kontrak)
}
