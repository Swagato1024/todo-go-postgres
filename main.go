package main

import (
	"fmt"
	"log"
	"os"

	"github.com/todo/db"
	"github.com/todo/app"
)

func main() {
	dbConnection, _ := db.ConnectDB()

	app := app.CreateApp(dbConnection)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}