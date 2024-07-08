package app

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/todo/db"
	"github.com/todo/models"
)

func TestSendTodos_Success(t *testing.T) {
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
	resp, errInReq := app.Test(req)

	assert.NoError(t, errInReq)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
}

func TestSendTodos_ServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := db.NewMockTodoRepository(ctrl)

	mockRepo.EXPECT().GetAllTodo().Return(nil, errors.New("mock error"))

	app := CreateApp(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	resp, _:= app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Expected Internal server Error")
}
