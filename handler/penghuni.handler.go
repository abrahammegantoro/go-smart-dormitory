package handler

import (
	"go-smart-dormitory/database"
	"go-smart-dormitory/model/dto"
	"go-smart-dormitory/model/entity"
	"log"
	"math/rand"
	"strconv"
	"time"
	"strings"
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
)

var (
	genderChoices = []entity.JenisKelamin{entity.LakiLaki, entity.Perempuan}
	statusChoices = []entity.Status{entity.Diterima, entity.MenungguPembayaran, entity.MenungguPembuatanKontrak, entity.BelumDireview}
)

func init() {
	rand.Seed(time.Now().UnixNano())
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

func CalonpenghuniHandlerRead(ctx *fiber.Ctx) error {
	var calonpenghuni []entity.Penghuni
	search := ctx.Query("search", "")
	
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		log.Println("Error parsing page parameter:", err)
		page = 1
	}

	limit := 10
	offset := (page - 1) * limit


	result := database.DB.Where("status != ? AND LOWER(nama) LIKE ?", "Diterima", "%"+strings.ToLower(search)+"%").Offset(offset).Limit(limit).Find(&calonpenghuni)

	if result.Error != nil {
		log.Println(result.Error)
	}

	var totalRecords int64
	database.DB.Model(&entity.Penghuni{}).Where("status != ? AND LOWER(nama) LIKE ?", "Diterima", "%"+strings.ToLower(search)+"%").Count(&totalRecords)

	totalPages := (totalRecords + int64(limit) - 1) / int64(limit)

	return ctx.JSON(fiber.Map{
		"data":       calonpenghuni,
		"totalPages": totalPages,
	})
}

func PenghuniHandlerRead(ctx *fiber.Ctx) error {
	search := ctx.Query("search", "")

	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil || page < 1 {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid page number")
	}

	pageSize, err := strconv.ParseFloat(ctx.Query("pageSize", strconv.Itoa(10)), 64)
	if err != nil || pageSize < 1 {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid page size")
	}

	var dtoResp = []dto.GetPenghuniAktifResponseDTO{}

	// Fetch Penghuni entities with related Kontrak entities where KamarID is not null
	result := database.DB.Table("penghunis").
		Select("kamars.nomor_kamar, penghunis.id, penghunis.nama, penghunis.jenis_kelamin, penghunis.nomor_telepon, penghunis.kontak_darurat").
		Joins("LEFT JOIN kontraks ON penghunis.id = kontraks.penghuni_id").
		Joins("LEFT JOIN kamars ON kontraks.kamar_id = kamars.id").
		Where("kontraks.kamar_id IS NOT NULL AND penghunis.status = ? AND LOWER(penghunis.nama) LIKE ?", "Diterima", "%"+strings.ToLower(search)+"%").
		Find(&dtoResp)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to fetch penghuni")
	}

	return ctx.JSON(
		fiber.Map{
			"data":      dtoResp,
			"page":      page,
			"totalPage": int(result.RowsAffected/int64(pageSize)) + 1,
		},
	)
}

func CalonPenghuniHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var penghuni entity.Penghuni
	result := database.DB.Where("id = ?", id).Delete(&penghuni)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to delete penghuni")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Penghuni deleted",
	})
}

func PenghuniHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var kontrak entity.Kontrak
	resultKontrak := database.DB.Where("penghuni_id = ?", id).First(&kontrak)

	if resultKontrak.Error != nil {
		log.Println(resultKontrak.Error)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to delete penghuni")
	}

	// change kamar status
	var kamar entity.Kamar
	resultKamar := database.DB.Where("id = ?", kontrak.KamarID).First(&kamar)
	if resultKamar.Error != nil {
		log.Println(resultKamar.Error)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to delete penghuni")
	}

	kamar.Status = "Available"
	database.DB.Save(&kamar)
	database.DB.Delete(&kontrak)

	var penghuni entity.Penghuni
	result := database.DB.Where("id = ?", id).Delete(&penghuni)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to delete penghuni")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Penghuni deleted",
	})
}

func PenghuniHandlerReadById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var calonpenghuni entity.Penghuni
	result := database.DB.Where("id = ?", id).First(&calonpenghuni)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to fetch calonpenghuni")
	}

	return ctx.JSON(fiber.Map{
		"data": calonpenghuni,
	})
}

func PenghuniHandlerUpdateStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var calonpenghuni entity.Penghuni
	result := database.DB.Where("id = ?", id).First(&calonpenghuni)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch calonpenghuni",
		})
	}

	if calonpenghuni.Status == "Menunggu Pembayaran" {
		calonpenghuni.Status = entity.Status("Diterima")
	} else if calonpenghuni.Status == "Belum Direview" {
		calonpenghuni.Status = entity.Status("Menunggu Pembuatan Kontrak")
	}

	result = database.DB.Save(&calonpenghuni)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch calonpenghuni",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Calonpenghuni updated",
		"data":    calonpenghuni,
	})
}
