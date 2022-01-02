package configs

import "github.com/gofiber/fiber/v2"

func Server() *fiber.Config {
	return &fiber.Config{
		AppName: "test",
	}
}
