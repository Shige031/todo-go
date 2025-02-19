package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Shige031/todo-go/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


func main() {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("USER_NAME")
	dbPass := os.Getenv("USER_PASS")
	dbName := os.Getenv("DB_NAME")

	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPass, dbName)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ルーターを作成する
	r := api.NewRouter(db)

	// サーバーを起動する
	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}