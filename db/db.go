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
	priority TEXT NOT NULL,
	completed INTEGER NOT NULL DEFAULT 0,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	DB.Exec(createToDosTable)

}

func GetRowCount() (int, error) {
	query := "SELECT COUNT(*) FROM todos"
	var count int

	err := DB.QueryRow(query).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
