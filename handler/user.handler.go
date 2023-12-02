package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-smart-dormitory/database"
	"github.com/bxcodec/faker/v3"
	"go-smart-dormitory/model/entity"
	"math/rand"
	"time"
	"strconv"
	"log"
)

var (
	genderChoices = []entity.JenisKelamin{entity.LakiLaki, entity.Perempuan}
	statusChoices = []entity.Status{entity.Diterima, entity.MenungguAlokasiKamar, entity.MenungguPembayaran, entity.MenungguPembuatanKontrak, entity.BelumDireview}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func UserHandlerRead(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World 👋!")
}

func PenghuniHandlerCreate(ctx *fiber.Ctx) error {
	for i := 0; i < 50; i++ {
		database.DB.Create(&entity.Penghuni{
			Email:                 faker.Email(),
			Nama:                  faker.Name(),
			NIM:                   strconv.FormatInt(faker.UnixTime(), 10), // Convert int64 to string
			JenisKelamin:          genderChoices[rand.Intn(len(genderChoices))],
			NomorTelepon:          faker.Phonenumber(),
			KontakDarurat:         faker.Phonenumber(),
			NamaKontakDarurat:     faker.Name(),
			HubunganKontakDarurat: faker.TitleFemale(),
			Alasan:                faker.Paragraph(),
			Status:                statusChoices[rand.Intn(len(statusChoices))],
		})
	}
	return ctx.SendString("Data Created")
}

const defaultPageSize = 10

func PenghuniHandlerRead(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page", "1"))

	if err != nil || page < 1 {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid page number")
	}

	pageSize, err := strconv.ParseFloat(ctx.Query("pageSize", strconv.Itoa(defaultPageSize)), 64)
	if err != nil || pageSize < 1 {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid page size")
	}

	var penghuni []entity.Penghuni
	result := database.DB.Where("status = ?", entity.Diterima).Offset((page - 1) * 10).Limit(10).Find(&penghuni)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to fetch penghuni")
	}

	return ctx.JSON(
		fiber.Map{
			"data": penghuni,
			"page": page,
			"totalPage": int(result.RowsAffected / int64(pageSize)) + 1,
		},
	)
}