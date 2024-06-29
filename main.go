package main

import (
	"fmt"

	"github.com/ash-tise/todo/db"
	"github.com/ash-tise/todo/todos"
)

func main() {

	db.InitDB()

	todo := todos.NewTodo("get this app done")

	err := todo.AddToDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("successfully added todo")
}
