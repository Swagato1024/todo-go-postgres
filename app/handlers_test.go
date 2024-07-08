package app

import (
	"encoding/json"
	"errors"
	"io"
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

	var actualTodos []models.Todo
    decoder := json.NewDecoder(resp.Body)
    if err := decoder.Decode(&actualTodos); err != nil {
        t.Fatalf("Failed to decode JSON response: %v", err)
    }

	assert.NoError(t, errInReq)
	assert.Equal(t, todoList, actualTodos)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "should send all todos")
}

func TestSendTodos_ServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := db.NewMockTodoRepository(ctrl)

	mockRepo.EXPECT().GetAllTodo().Return(nil, errors.New("mock error"))

	app := CreateApp(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	resp, _:= app.Test(req)

	body, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }

	assert.Equal(t, "Internal server error", string(body), "Response message mismatch")
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "should fail send todo for internal error")
}


func TestDeleteTodo_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := db.NewMockTodoRepository(ctrl)

	mockRepo.EXPECT().DeleteTodo("some-id").Return(nil)

	app := CreateApp(mockRepo)
	req := httptest.NewRequest(http.MethodDelete, "/todo/some-id", nil)
	resp, errInReq := app.Test(req)

	assert.NoError(t, errInReq)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "should delete todo by id")
}

func TestDeleteTodos_ServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := db.NewMockTodoRepository(ctrl)

	mockRepo.EXPECT().DeleteTodo("some-id").Return(errors.New("mock-error"))

	app := CreateApp(mockRepo)
	req := httptest.NewRequest(http.MethodDelete, "/todo/some-id", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "should fail todo deletion for internal error")
}
