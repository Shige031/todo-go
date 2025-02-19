package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Shige031/todo-go/models"
	"github.com/gorilla/mux"
)

func TestGetTodosHandler(t *testing.T) {
	
	t.Run("正常系", func(t *testing.T) {
		url := "http://localhost:8080/todos"
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rec := httptest.NewRecorder()

		todoController.GetTodosHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected %v, but got %v", http.StatusOK, rec.Code)
		}

		expectedLen := 5
		var todos []models.Todo
		if err := json.Unmarshal(rec.Body.Bytes(), &todos); err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}
		if len(todos) != expectedLen {
			t.Errorf("expected %v, but got %v", expectedLen, len(todos))
		}
	})
}

func TestGetTodoHandler(t *testing.T) {
	
	t.Run("正常系", func(t *testing.T) {
		url := "http://localhost:8080/todos/7595eb20-ee9a-11ef-9fe7-0242ac1c0002"
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rec := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/todos/{id}", todoController.GetTodoHandler).Methods(http.MethodGet)
		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected %v, but got %v", http.StatusOK, rec.Code)
		}

		var todo models.Todo
		if err := json.Unmarshal(rec.Body.Bytes(), &todo); err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}
		if todo.ID != "7595eb20-ee9a-11ef-9fe7-0242ac1c0002" {
			t.Errorf("expected %v, but got %v", "7595eb20-ee9a-11ef-9fe7-0242ac1c0002", todo.ID)
		}
	})
}

func TestCreateTodoHandler(t *testing.T) {
	
	t.Run("正常系", func(t *testing.T) {
		url := "http://localhost:8080/todos"

		// リクエストボディの作成
		todo := models.Todo{
			Title: "Test Todo",
		}
		body, err := json.Marshal(todo)
		if err != nil {
			t.Fatalf("failed to marshal request body: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		todoController.CreateTodoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected %v, but got %v", http.StatusOK, rec.Code)
		}
	})
}

func TestUpdateTodoHandler(t *testing.T) {
	
	t.Run("正常系", func(t *testing.T) {
		url := "http://localhost:8080/todos"

		// リクエストボディの作成
		todo := models.Todo{
			ID: "7595eb20-ee9a-11ef-9fe7-0242ac1c0002",
			Title: "Updated Todo",
		}
		body, err := json.Marshal(todo)
		if err != nil {
			t.Fatalf("failed to marshal request body: %v", err)
		}

		req := httptest.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		todoController.UpdateTodoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected %v, but got %v", http.StatusOK, rec.Code)
		}
	})
}

func TestDeleteTodoHandler(t *testing.T) {
	
	t.Run("正常系", func(t *testing.T) {
		url := "http://localhost:8080/todos/7595eb20-ee9a-11ef-9fe7-0242ac1c0002"
		req := httptest.NewRequest(http.MethodDelete, url, nil)
		rec := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/todos/{id}", todoController.DeleteTodoHandler).Methods(http.MethodDelete)
		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected %v, but got %v", http.StatusOK, rec.Code)
		}
	})
}