package main

import (
	"github.com/ash-tise/todo/db"
	"github.com/ash-tise/todo/todos"
)

func main() {

	db.InitDB()

	// todo := todos.NewTodo("test3")
	// err := todo.AddToDB()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err := todos.RemoveFromDB(1)

	// if err != nil {
	// 	return
	// }

	err := todos.DisplayTodos()

	if err != nil {
		return
	}

}
