package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routeHandler(app)

	app.Listen(":3000")
}

func routeHandler(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
