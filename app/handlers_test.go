package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/todo/db"
	"github.com/todo/models"
)

func TestSendTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := db.NewMockTodoRepository(ctrl)
	todoList := []models.Todo{
		{ID: "1", Title: "buy milk", Completed: false},
		{ID: "2", Title: "break for tea", Completed: false},
	}

	mockRepo.EXPECT().GetAllTodo().Return(todoList, nil)

	app := CreateApp(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
}
