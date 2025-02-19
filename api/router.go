package api

import (
	"database/sql"

	"github.com/Shige031/todo-go/controllers"
	"github.com/Shige031/todo-go/usecases"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	usecases := usecases.NewUsecases(db)
	todoController := controllers.NewTodoController(usecases)

	r := mux.NewRouter()

	r.HandleFunc("/todos", todoController.GetTodosHandler).Methods("GET")
	r.HandleFunc("/todos/{id}", todoController.GetTodoHandler).Methods("GET")
	r.HandleFunc("/todos", todoController.CreateTodoHandler).Methods("POST")
	r.HandleFunc("/todos", todoController.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/todos/{id}", todoController.DeleteTodoHandler).Methods("DELETE")

	return r
}