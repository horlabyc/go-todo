package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	apiRoute := app.Group("/api")

	//Test Handler
	apiRoute.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome on board")
	})

	log.Fatal(app.Listen(":4000"))
}
