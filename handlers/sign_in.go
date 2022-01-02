package handlers

import (
	"ecommerce-v1/auth/models"
	"ecommerce-v1/auth/services"
	"ecommerce-v1/auth/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) error {
	var model models.SignIn
	c.BodyParser(&model)

	account, err := services.SignIn(&model)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.HttpResponse{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.HttpResponse{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{Value: refreshToken, Name: "refresh_token", Path: "/", Secure: true, HTTPOnly: true, SameSite: "lax"})

	return c.Status(http.StatusOK).JSON(models.HttpResponse{
		Success: true,
		Data:    account,
		Error:   nil,
	})
}
