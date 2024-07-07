package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func reqlogger(c *fiber.Ctx) error {
	log.Printf("Request: %s %s",c.Method(), c.Path())

	if err := c.Next(); err != nil {
		return err
	}

	return nil
}

func CreateApp() *fiber.App {
	app := fiber.New()

	app.Use(reqlogger)

	app.Get("/", IndexHandler)
	app.Post("/post", AddTodo)
	app.Delete("/delete", DeleteTodo)

	return app
}