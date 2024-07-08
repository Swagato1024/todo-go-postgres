package app

import (
	"fmt"
	"log"

	"github.com/todo/db"
	"github.com/todo/models"

	"github.com/gofiber/fiber/v2"
)

func reqlogger(c *fiber.Ctx) error {
	log.Printf("Request: %s %s", c.Method(), c.Path())

	if err := c.Next(); err != nil {
		return err
	}

	return nil
}

func CreateApp(r db.TodoRepository) *fiber.App {
	app := fiber.New()

	app.Use(reqlogger)

	app.Get("/todos", func(c *fiber.Ctx) error {
		fmt.Printf("Inside get todos");
		
		sendTodos(c, r)
		return nil
	})

	app.Post("/todo", func(c *fiber.Ctx) error {
		var todo models.Todo

		if err := c.BodyParser(&todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON payload",
			})
		}

		addTodo(c, r, todo)
		return nil
	})

	app.Delete("/delete/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		deleteTodo(c, r, id)
		return nil
	})

	return app
}
