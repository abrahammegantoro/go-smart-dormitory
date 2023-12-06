package handler

import (
	"go-smart-dormitory/database"
	"go-smart-dormitory/model/dto"
	"go-smart-dormitory/model/entity"
	"go-smart-dormitory/utils"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func AdminHandlerLogin(ctx *fiber.Ctx) error {
	loginRequest := new(dto.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed",
			"error":   err.Error(),
		})
	}

	// validate request
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed",
			"error":   errValidate.Error(),
		})
	}

	var admin entity.Admin
	err := database.DB.Where("username = ?", loginRequest.Username).First(&admin).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong username or password",
		})
	}

	// check password
	isValid := utils.CheckPassword(admin.Password, loginRequest.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong username or password",
		})
	}

	// Generate token
	claims := jwt.MapClaims{}
	claims["username"] = admin.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := utils.GenerateToken(&claims)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
