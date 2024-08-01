package todo

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	service TodoInterface
}

func NewTodoController() *TodoController {
	return &TodoController{NewTodoService()}
}

func (controller *TodoController) CreateTodo(c *fiber.Ctx) error {
	var todo Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := controller.service.CreateTodo(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(todo)
}

func (controller *TodoController) GetList(c *fiber.Ctx) error {
	todos := controller.service.GetList()
	return c.JSON(todos)
}

func (controller *TodoController) UpdateTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	todo, err := controller.service.UpdateTodo(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(todo)
}

func (controller *TodoController) DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	if err := controller.service.DeleteTodo(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
