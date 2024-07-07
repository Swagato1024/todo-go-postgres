package main

import (
	"fmt"
	"log"
	"os"

	"github.com/todo/app"
)

func main() {
	app := app.CreateApp()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}