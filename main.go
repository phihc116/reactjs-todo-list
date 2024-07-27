package main

import (
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json: "id"`
	Completed bool   `json: "completed"`
	Body      string `json: "body"`
}

func main() {
	app := fiber.New()

	controller := NewTodoController()
	controller.RegisterRoutes(app)

	app.Listen(":8001")
}
