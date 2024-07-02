package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	homeDir, err := os.UserHomeDir()
	dbPath := homeDir + "/Todos/todos.db"

	if err != nil {
		log.Fatal("Failed to fetch User's Home directory")
	}

	DB, err = sql.Open("sqlite3", dbPath)

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
