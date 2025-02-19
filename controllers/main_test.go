package controllers_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Shige031/todo-go/controllers"
	"github.com/Shige031/todo-go/usecases"
	_ "github.com/go-sql-driver/mysql"
)

var todoController *controllers.TodoController

func TestMain(m *testing.M){
	dbUser := "user"
	dbPassword := "pass"
	dbDatabase := "todo-go-db-test"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}

	usecases := usecases.NewUsecases(db)
	todoController = controllers.NewTodoController(usecases)

	setUpTestData(db)

	m.Run()

	cleanUpTestData(db)
}

func setUpTestData(db *sql.DB) {
	_, err := db.Exec("INSERT INTO todos (id, title, completed) VALUES " +
  "('7595eb20-ee9a-11ef-9fe7-0242ac1c0002', 'Buy groceries', FALSE), " +
  "(UUID(), 'Finish Go project', TRUE), " +
  "(UUID(), 'Read a book', FALSE), " +
  "(UUID(), 'Exercise for 30 minutes', FALSE), " +
  "(UUID(), 'Call mom', TRUE);")
	if err != nil {
		fmt.Println(err)
	}
}

func cleanUpTestData(db *sql.DB) {
	_, err := db.Exec("DELETE FROM todos;")
	if err != nil {
		fmt.Println(err)
	}
}