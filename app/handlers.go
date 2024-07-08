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

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todos": todoList,
	})
}

func AddTodo(c *fiber.Ctx, r db.TodoRepository, todo models.Todo) {
	err := r.AddTodo(todo)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Todo creation failed",
		})
		return
	}

	c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo created successfully",
	})
}

func DeleteTodo(c *fiber.Ctx, r db.TodoRepository, id string) {
	err := r.DeleteTodo(id)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Todo creation failed",
		})
		return
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todos": "Todo deleted",
	})
}
