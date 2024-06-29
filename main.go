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

	todos.DisplayTodos()

}
