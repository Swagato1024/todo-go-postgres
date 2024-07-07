package app

import (
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	c.SendString("Welcome To Home!!")
	
	return nil
}

func AddTodo(c *fiber.Ctx) error {
	c.SendString("successfully created resource")
	return nil
}

func DeleteTodo(c *fiber.Ctx) error {
	c.SendString("Delete todo")
	return nil
}