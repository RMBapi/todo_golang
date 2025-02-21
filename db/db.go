package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB() {

	DB, err = sql.Open("sqlite3", "todo.db")

	if err != nil {
		panic("Couldn't connect database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {

	createTodoTable := `
	CREATE TABLE IF NOT EXISTS todoList (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task TEXT NOT NULL,
	description TEXT NOT NULL,
    datetime TEXT NOT NULL,
	status TEXT NOT NULL
    )
	`
	_, err := DB.Exec(createTodoTable)

	if err != nil {
		panic("Couldn't connect event table")
	}

}
