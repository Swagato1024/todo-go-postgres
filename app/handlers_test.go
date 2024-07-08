package app

import (
	"encoding/json"
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

	var actualTodos []models.Todo
    decoder := json.NewDecoder(resp.Body)
    if err := decoder.Decode(&actualTodos); err != nil {
        t.Fatalf("Failed to decode JSON response: %v", err)
    }

	assert.NoError(t, errInReq)
	assert.Equal(t, todoList, actualTodos)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Todos sent successfully")
}

func TestSendTodos_ServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := db.NewMockTodoRepository(ctrl)

	mockRepo.EXPECT().GetAllTodo().Return(nil, errors.New("mock error"))

	app := CreateApp(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	resp, _:= app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Send Todo Expected Internal server Error")
}


func TestDeleteTodo_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := db.NewMockTodoRepository(ctrl)

	mockRepo.EXPECT().DeleteTodo("some-id").Return(nil)

	app := CreateApp(mockRepo)
	req := httptest.NewRequest(http.MethodDelete, "/todo/some-id", nil)
	resp, errInReq := app.Test(req)

	// var responseMsg string
    // decoder := json.NewDecoder(resp.Body)
    // if err := decoder.Decode(&responseMsg); err != nil {
    //     t.Fatalf("Failed to decode JSON response: %v", err)
    // }

	assert.NoError(t, errInReq)
	// assert.Equal(t, "Todo deleted",responseMsg)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Delete todo by id successful")
}

func TestDeleteTodos_ServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := db.NewMockTodoRepository(ctrl)

	mockRepo.EXPECT().GetAllTodo().Return(nil, errors.New("mock error"))

	app := CreateApp(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	resp, _:= app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Expected Internal server Error")
}
