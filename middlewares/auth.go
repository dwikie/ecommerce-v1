package middlewares

import (
	"ecommerce-v1/auth/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Cookies("refresh_token")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(models.HttpResponse{
			Error:   "Unauthorized",
			Success: false,
			Data:    nil,
		})
	}
	return c.Next()
}
