package routes

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	handlers "superb.com/v0/user_service/handlers"
)

func InitRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	baseURL := os.Getenv("USER_SERVICE_BASE")
	log.Print("Base URL: " + baseURL)
	userApi := app.Group(baseURL)
	userApi.Post("/", handlers.CreateUser)
	userApi.Get("/:id", handlers.GetUser)
}