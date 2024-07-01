package todos

import (
	"errors"
	"fmt"
	"log"

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

func RemoveFromDB(index int) error {

	todos, err := fetchRows()

	if err != nil {
		return err
	}

	if len(todos) == 0 {
		fmt.Println("You currently have no ToDos!")
		return errors.New("no todos to remove")
	}

	for i, todo := range todos {
		if i+1 == index {
			key := todo.Task
			query := "DELETE FROM todos WHERE todo = ?"
			_, err := db.DB.Exec(query, key)

			if err != nil {
				log.Fatal(err)
			}

			return nil
		}
	}

	return errors.New("could not find index")

}

func fetchRows() ([]Todo, error) {

	query := `SELECT todo, priority FROM todos`
	var todos []Todo

	rows, err := db.DB.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var todo Todo

		err := rows.Scan(&todo.Task, &todo.Priority)

		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)

	}

	return todos, nil

}

func DisplayTodos() error {

	todos, err := fetchRows()

	if err != nil {
		return err
	}

	if len(todos) == 0 {
		fmt.Println("You currently have no Todos!")
		return nil
	}

	for index, todo := range todos {
		fmt.Printf("%d)  %s  |  %s\n", index+1, todo.Task, todo.Priority)
	}

	return nil

}
