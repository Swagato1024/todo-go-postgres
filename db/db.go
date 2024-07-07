package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error){
	// connectionStr := "user=postgres password=secret dbname=postgres sslmode=disable"

	cfg := getConfig()

    connectionInfo := fmt.Sprintf(
		"host=%s port=%s user=%s " +
        "password=%s dbname=%s sslmode=disable",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	conn, err := sql.Open("postgres", connectionInfo)

	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		panic(err)
	}

	return conn, nil


	// rows, err := conn.Query("SELECT version();")
	// if err != nil {
	// 	panic(err)
	// }

	// for rows.Next() {
	// 	var version string
	// 	rows.Scan(&version)
	// 	fmt.Println(version)
	// }

	// rows.Close()

	// Create a new database
	// _, err = conn.Exec("CREATE DATABASE employee TEMPLATE template0")
	// Create a new database from template0

	// if err != nil {
	// 	log.Fatal("Error creating database:", err)
	// }

	    // Create the 'users' table if it does not exist
	// 	createTable := `
    //     CREATE TABLE IF NOT EXISTS users (
    //         id SERIAL PRIMARY KEY,
    //         username VARCHAR(50) NOT NULL,
    //         email VARCHAR(100) NOT NULL UNIQUE
    //     )
    // `

    // _, err = conn.Exec(createTable)
	// if err != nil {
    //     log.Fatal("Error creating table:", err)
    // }


	// Example SQL insert statement
	// insertStatement := `
	// 	 INSERT INTO users (username, email)
	// 	 VALUES ($1, $2)
	//  `

	// Example data to insert
	// username := "john_doe"
	// email := "john.doe@example.com"

	// Execute the SQL statement
	// _, err = conn.Exec(insertStatement, username, email)
	// if err != nil {
	// 	log.Fatal("Error inserting data into table:", err)
	// }

	// fmt.Println("Data inserted successfully!")

	// conn.Close()
}
