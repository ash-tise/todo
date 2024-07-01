package main

import (
	"github.com/ash-tise/todo/db"
	"github.com/ash-tise/todo/todos"
)

func main() {

	db.InitDB()

	todos.DisplayTodos()

}
