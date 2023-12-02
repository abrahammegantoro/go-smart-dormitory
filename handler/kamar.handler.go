package handler

import (
	"go-smart-dormitory/database"
	"go-smart-dormitory/model/entity"
	"math/rand"

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
