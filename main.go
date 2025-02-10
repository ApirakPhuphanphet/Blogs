package main

import (
	"github.com/ApirakPhuphanphet/Go-Blogs/db"
	"github.com/ApirakPhuphanphet/Go-Blogs/model"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database := db.DatabaseConnection()
	database.AutoMigrate(&model.User{})

	app := fiber.New()

	routeHandler(app)

	app.Listen(":3000")
}

func routeHandler(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
