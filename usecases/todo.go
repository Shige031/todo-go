package usecases

import (
	"github.com/Shige031/todo-go/models"
	"github.com/Shige031/todo-go/repositories"
)

type TodoUsecase interface {
	GetTodos() ([]models.Todo, error)
	GetTodo(id string) (models.Todo, error)
	CreateTodo(todo models.Todo) (models.Todo, error)
	UpdateTodo(todo models.Todo) (models.Todo, error)
	DeleteTodo(id string) error
} 

func (u *Usecases) GetTodos()([]models.Todo, error){
		todos, err := repositories.GetTodos(u.db)
		if err != nil {
			return nil, err
		}
		return todos, nil
}

func (u *Usecases) GetTodo(id string)(models.Todo, error){
	todo, err := repositories.GetTodo(u.db, id)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (u *Usecases) CreateTodo(todo models.Todo)(models.Todo, error){
	todo, err := repositories.CreateTodo(u.db, todo)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (u *Usecases) UpdateTodo(todo models.Todo)(models.Todo, error){
	todo, err := repositories.UpdateTodo(u.db, todo)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (u *Usecases) DeleteTodo(id string) error {
	err := repositories.DeleteTodo(u.db, id)
	if err != nil {
		return err
	}
	return nil
}