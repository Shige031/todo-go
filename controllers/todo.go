package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Shige031/todo-go/models"
	"github.com/Shige031/todo-go/usecases"
	"github.com/gorilla/mux"
)

type TodoController struct {
	usecase usecases.TodoUsecase
}

func NewTodoController(u usecases.TodoUsecase) *TodoController {
	return &TodoController{usecase: u}
}

func (c *TodoController) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := c.usecase.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos);
}

func (c *TodoController) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	todo, err := c.usecase.GetTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo);
}

func (c *TodoController) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var requestTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&requestTodo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	todo, err := c.usecase.CreateTodo(requestTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo);
}

func (c *TodoController) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var requestTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&requestTodo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	todo, err := c.usecase.UpdateTodo(requestTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo);
}

func (c *TodoController) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := c.usecase.DeleteTodo(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}