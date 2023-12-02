package handler

import (
	"go-smart-dormitory/database"
	"go-smart-dormitory/model/entity"
	"log"
	"math/rand"
	"strconv"
	"time"
	"go-smart-dormitory/model/dto"
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func UserHandlerRead(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World 👋!")
}

func CalonpenghuniHandlerRead(ctx *fiber.Ctx) error {
	var calonpenghuni []entity.Penghuni

	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		log.Println("Error parsing page parameter:", err)
		page = 1
	}

	limit := 10
	offset := (page - 1) * limit

	result := database.DB.Where("status != ?", "Diterima").Offset(offset).Limit(limit).Find(&calonpenghuni)

	if result.Error != nil {
		log.Println(result.Error)
	}

	var totalRecords int64
	database.DB.Model(&entity.Penghuni{}).Where("status != ?", "Diterima").Count(&totalRecords)

	totalPages := (totalRecords + int64(limit) - 1) / int64(limit)

	return ctx.JSON(fiber.Map{
		"data":       calonpenghuni,
		"totalPages": totalPages,
	})
}

func PenghuniHandlerRead(ctx *fiber.Ctx) error {
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
    	Select("kamars.nomor_kamar, penghunis.nama, penghunis.jenis_kelamin, penghunis.nomor_telepon, penghunis.kontak_darurat").
    	Joins("LEFT JOIN kontraks ON penghunis.id = kontraks.penghuni_id").
    	Joins("LEFT JOIN kamars ON kontraks.kamar_id = kamars.id").
    	Where("kontraks.kamar_id IS NOT NULL").
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

func PenghuniHandlerDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var penghuni entity.Penghuni
	result := database.DB.Where("id = ?", id).Delete(&penghuni)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to delete penghuni")
	}

	return ctx.SendString("Penghuni successfully deleted")
}

func PenghuniHandlerReadById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var calonpenghuni entity.Penghuni
	result := database.DB.Where("id = ?", id).First(&calonpenghuni)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to fetch calonpenghuni")
	}

	return ctx.JSON(calonpenghuni)
}