package todo

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	resultId, err := controller.service.CreateTodo(todo)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(resultId)
}

func (controller *TodoController) GetList(c *fiber.Ctx) error {
	todos, err := controller.service.GetList()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(todos)
}

func (controller *TodoController) UpdateTodo(c *fiber.Ctx) error {
	var todo TodoRequestUpdate
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := controller.service.UpdateTodo(todo)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func (controller *TodoController) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	result, err := controller.service.DeleteTodo(objectId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}
