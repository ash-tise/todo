package todos

import (
	"github.com/ash-tise/todo/db"
)

type Todo struct {
	Task     string
	Priority string
}

func NewTodo(task string) Todo {
	return Todo{
		Task:     task,
		Priority: "medium",
	}
}

func NewTodoWithPriority(task, priority string) Todo {
	return Todo{
		Task:     task,
		Priority: priority,
	}
}

func (t Todo) AddToDB() error {
	query := `INSERT INTO todos (todo, priority) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(t.Task, t.Priority)

	if err != nil {
		return err
	}

	return nil
}
