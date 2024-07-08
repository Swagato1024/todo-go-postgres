package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/todo/db"
	"github.com/todo/models"
)

func sendTodos(c *fiber.Ctx, r db.TodoRepository) {
	todoList, err := r.GetAllTodo()

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server error",
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(todoList)
}

func addTodo(c *fiber.Ctx, r db.TodoRepository, todo models.Todo) {
	err := r.AddTodo(todo)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Todo creation Failed")
		return
	}

	c.Status(fiber.StatusCreated).SendString("Todo added successfully")
}

func deleteTodo(c *fiber.Ctx, r db.TodoRepository, id string) {
	err := r.DeleteTodo(id)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Todo deletion failed")
		return
	}
	c.Status(fiber.StatusOK).SendString("Todo deleted")
}
