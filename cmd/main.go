package main

import (
	"synapsehr/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Static("/home", "./public")

	app.Get("/user", routes.GetUser)
	app.Post("/user", routes.CreateUser)
	app.Get("/", routes.Home)
	app.Listen(":8080")
}
