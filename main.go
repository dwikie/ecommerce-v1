package main

import (
	"ecommerce-v1/auth/configs"
	"ecommerce-v1/auth/handlers"
	"ecommerce-v1/auth/repository"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	configs.InitDB()

	app := fiber.New(*configs.Server())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	router(app)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}

func router(app *fiber.App) {
	app.Post("/sign-up", handlers.SignUp)

	app.Post("/sign-in", handlers.SignIn)

	app.Get("/sign-out", func(c *fiber.Ctx) error {
		c.ClearCookie("refresh_token")
		return c.SendString("Sign Out")
	})

	app.Get("/account", func(c *fiber.Ctx) error {
		email := c.Query("email", "")
		data, err := repository.Account.GetAccountByEmail(email)
		if err != nil {
			return c.JSON(err.Error())
		}

		// configs.DB.Offset(0).Limit(2).Find(&account).Where("updated_at IS NOT NULL")

		return c.JSON(data)
	})
}
