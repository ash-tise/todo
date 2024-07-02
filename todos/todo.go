package todos

import (
	"errors"
	"fmt"
	"log"

	"github.com/ash-tise/todo/db"
)

type Todo struct {
	Task      string
	Priority  string
	Completed bool
}

func NewTodo(task string) Todo {
	return Todo{
		Task:      task,
		Priority:  "medium",
		Completed: false,
	}
}

func NewTodoWithPriority(task, priority string) Todo {
	return Todo{
		Task:      task,
		Priority:  priority,
		Completed: false,
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

	fmt.Println("Could not find index!")
	return errors.New("could not find index")

}

func fetchRows() ([]Todo, error) {

	query := `SELECT todo, priority FROM todos ORDER BY created_at`
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

	green := "\033[32m"
	reset := "\033[0m"

	todos, err := fetchRows()

	if err != nil {
		return err
	}

	if len(todos) == 0 {
		fmt.Print("You currently have no Todos!\n\n")
		return nil
	}

	fmt.Print("\n")
	for index, todo := range todos {
		// coudln't get dynamic padding to work
		// padding := width - len(todo.Task)

		if todo.checkIfCompleted() {
			// fmt.Print(strconv.FormatInt(int64(index+1), 10)+") "+green+todo.Task+reset+strings.Repeat(" ", padding), "\n"+strings.Repeat(" ", width+6), "\n")
			fmt.Printf("%d)  %s%s    completed%s\n\n", index+1, green, todo.Task, reset)
		} else {
			fmt.Printf("%d)  %s\n\n", index+1, todo.Task)
			// fmt.Print(strconv.FormatInt(int64(index+1), 10)+") "+todo.Task+strings.Repeat(" ", padding), "\n"+strings.Repeat(" ", width+6), "\n")
			// fmt.Printf("%d)  %-*s  |  %s\n", index+1, padding, todo.Task, todo.Priority)

		}
	}

	return nil

}

func getLongestTodoLen() int {
	todos, err := fetchRows()

	if err != nil {
		return 0
	}

	var max int

	for _, todo := range todos {
		if len(todo.Task) > max {
			max = len(todo.Task)
		}
	}

	return max
}

func MarkAsCompleted(index int) error {
	todos, err := fetchRows()

	if err != nil {
		return err
	}

	for i, todo := range todos {
		if i+1 == index {
			query := "UPDATE todos SET completed = 1 WHERE todo = ?"
			_, err := db.DB.Exec(query, todo.Task)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t Todo) checkIfCompleted() bool {
	query := "SELECT completed FROM todos WHERE todo = ?"

	var val int

	err := db.DB.QueryRow(query, t.Task).Scan(&val)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return val == 1
}
