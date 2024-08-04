package todo

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *TodoController) RegisterRoutes(app *fiber.App) {
	app.Post("api/todos", controller.CreateTodo)
	app.Get("api/todos", controller.GetList)
	app.Put("api/todos", controller.UpdateTodo)
	app.Delete("api/todos/:id", controller.DeleteTodo)
}
