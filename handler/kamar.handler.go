package handler

import (
	"go-smart-dormitory/database"
	"go-smart-dormitory/model/entity"
	"log"
	"math/rand"
	"strconv"

	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
)

var (
	statusKamarChoices = []entity.StatusKamar{entity.Available, entity.Booked, entity.Occupied}
)

func KamarHandlerCreate(ctx *fiber.Ctx) error {
	for i := 0; i < 50; i++ {
		database.DB.Create(&entity.Kamar{
			NomorKamar: uint8(i + 1),
			Fasilitas:  faker.Sentence(),
			Status:     statusKamarChoices[rand.Intn(len(statusKamarChoices))],
		})
	}
	return ctx.SendString("Data Created")
}

func GetKamarHandlerRead(ctx *fiber.Ctx) error {
	// Pagination
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil || page < 1 {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid page number",
		})
	}

	pageSize, err := strconv.Atoi(ctx.Query("pageSize", "10"))
	if err != nil || pageSize < 1 {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid page size",
		})
	}

	// Fetch paginated data from the database
	var kamar []entity.Kamar
	result := database.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&kamar)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	var totalRecords int64
	database.DB.Model(&entity.Kamar{}).Where("status != ?", "Diterima").Count(&totalRecords)
	limit := 10
	totalPages := (totalRecords + int64(limit) - 1) / int64(limit)

	return ctx.JSON(fiber.Map{
		"data":      kamar,
		"page":      page,
		"totalPage": totalPages,
	})
}

func UpdateStatusKamarHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	status := ctx.Query("status")

	var kamar entity.Kamar
	result := database.DB.First(&kamar, id)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	kamar.Status = entity.StatusKamar(status)
	err := database.DB.Save(&kamar)

	if err.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Data updated",
		"data":    kamar,
	})
}

func KamarAvailableHandleRead(ctx *fiber.Ctx) error {
	var kamar []entity.Kamar

	result := database.DB.Where("status = ?", entity.Available).Find(&kamar)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return ctx.JSON(kamar)
}
