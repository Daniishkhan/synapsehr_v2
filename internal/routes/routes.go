package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/user", GetUser)
	app.Post("/user", CreateUser)
	app.Get("/", Home)
	app.Post("/greet", Greet)
}
