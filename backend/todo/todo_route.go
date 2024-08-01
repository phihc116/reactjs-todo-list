package todo

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *TodoController) RegisterRoutes(app *fiber.App) {
	app.Post("/todos", controller.CreateTodo)
	app.Get("/todos", controller.GetList)
	app.Put("/todos/:id", controller.UpdateTodo)
	app.Delete("/todos/:id", controller.DeleteTodo)
}
