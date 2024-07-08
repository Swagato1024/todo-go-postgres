package db

import (
	"github.com/todo/models"
)

type TodoRepository interface {
	GetAllTodo() ([]models.Todo, error)
	AddTodo(todo models.Todo) error
	DeleteTodo(id string) error
}