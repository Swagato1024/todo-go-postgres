package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/todo/models"
)

type DB struct {
	Connection *sql.DB
}

func (db *DB) GetAllTodo() ([]models.Todo, error) {
	rows, err := db.Connection.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, fmt.Errorf("error querying todos: %v", err)
	}

	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, fmt.Errorf("error scanning todo: %v", err)
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating over rows: %v", err)
    }

	return todos, nil
}

func (db *DB) AddTodo(todo models.Todo) error {
	_, err := db.Connection.Exec("INSERT INTO todos (id, title, completed) VALUES ($1, $2, $3)",
		todo.ID, todo.Title, todo.Completed)

	if err != nil {
		return  fmt.Errorf("error inserting todo: %v", err)
	}

	return nil
}

func (db *DB) DeleteTodo(id string) error {
	_, err := db.Connection.Exec("DELETE FROM todos WHERE id = $1", id)

	if err != nil {
		return fmt.Errorf("error deleting todo: %v", err)
	}
	return nil
}

func createTable(db *sql.DB, tableName string) error {
	createTableQuery := fmt.Sprintf(`
	    CREATE TABLE IF NOT EXISTS %s (
	        id SERIAL PRIMARY KEY,
	        title VARCHAR(50) NOT NULL,
	        completed BOOLEAN
	    )
	`, tableName)

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating table %s: %v", tableName, err)
	}

	return nil
}

func ConnectDB() (*DB, error) {
	cfg := getConfig()

	connectionInfo := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	conn, err := sql.Open("postgres", connectionInfo)

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	if err := conn.Ping(); err != nil {
        return nil, fmt.Errorf("error pinging database: %v", err)
    }


    tableName := "todos"

    query := "DROP TABLE " + tableName

    conn.Exec(query)

	err = createTable(conn, "todos")

	if err != nil {
		return nil, fmt.Errorf("error while creating the table")
	}

	return &DB{Connection: conn}, nil
}

