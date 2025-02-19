package repositories

import (
	"database/sql"

	"github.com/Shige031/todo-go/models"
)

func GetTodos(db * sql.DB)([]models.Todo, error){
	const query = `SELECT * FROM todos`

	todos := []models.Todo{}
	rows, err := db.Query(query)
	if err != nil {
		return todos, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func GetTodo(db * sql.DB, id string)(models.Todo, error){
	const query = `SELECT * FROM todos WHERE id = ?`

	var todo models.Todo
	if err := db.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		return todo, err
	}

	return todo, nil
}

func CreateTodo(db *sql.DB, todo models.Todo)(models.Todo, error) {
	const query = `INSERT INTO todos ( title ) VALUES (?)`

	_, err := db.Exec(query, todo.Title)
	if err != nil {	
		return todo, err
	}

	return todo, nil
}

func UpdateTodo(db *sql.DB, todo models.Todo)(models.Todo, error) {
	const query = `UPDATE todos SET title = ?, completed = ? WHERE id = ?`

	_, err := db.Exec(query, todo.Title, todo.Completed, todo.ID)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func DeleteTodo(db *sql.DB, id string) error {
	const query = `DELETE FROM todos WHERE id = ?`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}