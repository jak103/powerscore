package server

import (
	"powerscore-server/internal/middleware"
	"powerscore-server/internal/server/apis"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler:          globalErrorHandler,
		DisableStartupMessage: true,
	})

	middleware.Setup(app)

	apis.SetupRoutes(app)

	app.Static("/", "../built/dist")

	app.Listen(":8000")
}

func globalErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).SendString("Not found")
}
