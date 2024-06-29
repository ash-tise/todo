package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "todos.db")

	if err != nil {
		panic("Failed to open database!")
	}

	createToDosTable := `
	CREATE TABLE IF NOT EXISTS todos (
	todo TEXT PRIMARY KEY,
	priority TEXT NOT NULL
	)
	`

	DB.Exec(createToDosTable)

}
