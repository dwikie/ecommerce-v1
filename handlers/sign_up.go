package handlers

import (
	"ecommerce-v1/auth/models"
	"ecommerce-v1/auth/services"
	"ecommerce-v1/auth/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	var model models.SignUp

	c.BodyParser(&model)

	validate := utils.StructValidator(&model)
	if validate != nil {
		return c.Status(http.StatusBadRequest).JSON(models.HttpResponse{
			Success: false,
			Data:    nil,
			Error:   validate,
		})
	}

	data, err := services.SignUp(&model)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(models.HttpResponse{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(models.HttpResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	})
}
