package routes

import (
	"synapsehr/views"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	return c.SendString("Hello, User!")
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}
	c.Response().Header.Set("Content-Type", "application/json")
	return c.JSON(user)
}

func Home(c *fiber.Ctx) error {
	//how do i send html here
	c.Set("Content-Type", "text/html")
	return views.Hello("Danish").Render(c.Context(), c.Response().BodyWriter())
}
