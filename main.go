package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/horlabyc/go-todo/database"
	"github.com/horlabyc/go-todo/todo"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	fmt.Println("heree")
	defer database.DB.Close()
	apiRoute := app.Group("/api")
	todo.Register(apiRoute, database.DB)

	log.Fatal(app.Listen(":5000"))
}
