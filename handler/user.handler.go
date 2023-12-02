package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-smart-dormitory/database"
	"github.com/bxcodec/faker/v3"
	"go-smart-dormitory/model/entity"
	"math/rand"
	"time"
	"strconv"
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