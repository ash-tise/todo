package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ash-tise/todo/db"
	"github.com/ash-tise/todo/todos"
)

const (
	reset  = "\033[0m"
	orange = "\033[38;5;214m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

func main() {

	if len(os.Args) < 2 {

		fmt.Printf("\nWelcome to the Todo CLI tool!\n------------------------------\n\n%sAdd todo:%s todo [ add | -a ] <task>\n\n%sRemove todo:%s todo [ remove | -r ] <task_id>\n\n%sList todos:%s todo [list | -l ]\n\n%sMark todo as completed:%s todo [ complete | -c ] <task_id>\n\n", orange, reset, red, reset, yellow, reset, green, reset)

		fmt.Print("Usage: todo <command> [arguments]\n\n")
		return
	}

	db.InitDB()

	command := os.Args[1]

	switch command {
	case "complete", "-c":
		if len(os.Args) < 3 {
			fmt.Print("Usage: todo [ complete | -c ] <task_id>\n\n")
			return
		}
		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid task ID")
			fmt.Print("Usage: todo [ complete | -c ] <task_id>\n\n")
			return
		}

		err = todos.MarkAsCompleted(id)

		if err != nil {
			fmt.Println(err)
			return
		}
	case "-a", "add":
		if len(os.Args) < 3 {
			fmt.Print("Usage: todo [ add | -a ] <task> \n\n")
			return
		}
		taskName := os.Args[2]
		err := todos.NewTodo(taskName).AddToDB()

		if err != nil {
			fmt.Println(err)
			fmt.Println("\nUsage: todo [ add | -a ] <task>")
			return
		}
	case "remove", "-r":
		if len(os.Args) < 3 {
			fmt.Print("Usage: todo [remove | -r ] <task_id>\n\n")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid task ID")
			fmt.Print("Usage: todo [ remove | -r ] <task_id>\n\n")
			return
		}

		err = todos.RemoveFromDB(id)

		if err != nil {
			fmt.Println(err)
			fmt.Print("Usage: todo [ remove | -r ] <task_id>\n\n")
			return
		}
	case "list", "-l":
		err := todos.DisplayTodos()

		if err != nil {
			fmt.Println(err)
			return
		}
	case "clear":
		if _, err := os.Stat("todos.db"); os.IsNotExist(err) {
			fmt.Printf("Your ToDo list is already cleared!")
			return
		}

		err := os.Remove("todos.db")
		if err != nil {
			fmt.Printf("Failed to clear ToDos")
			return
		}
	case "help", "-h":
		fmt.Printf("\nWelcome to the Todo CLI tool!\n------------------------------\n\n%sAdd todo:%s todo [ add | -a ] <task>\n\n%sRemove todo:%s todo [ remove | -r ] <task_id>\n\n%sList todos:%s todo [list | -l ]\n\n%sMark todo as completed:%s todo [ complete | -c ] <task_id>\n\n", orange, reset, red, reset, yellow, reset, green, reset)

		fmt.Print("Usage: todo <command> [arguments]\n\n")
		return

	default:
		fmt.Println("Unknown command:", command)
		fmt.Print("Usage: todo <command> [arguments]\n\n")
		return

	}
}
