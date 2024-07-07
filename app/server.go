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

func indexHandler(c *fiber.Ctx) error {
	c.SendString("Welcome To Home!!")
	
	return nil
}

func CreateApp() *fiber.App {
	app := fiber.New()

	app.Use(reqlogger)

	app.Get("/", indexHandler)
	// app.Post("/", postHandler)
	// app.Put("/update", putHandler)
	// app.Delete("/delete", deleteHandler)

	return app
}