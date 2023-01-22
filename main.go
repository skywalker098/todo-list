package main

import (
	"fmt"

	"github.com/todolist_iteration1/todo"
)

func main() {
	NewTodoServiceIns := todo.NewTodoService()
	//add
	fmt.Println(NewTodoServiceIns.AddTodo("Argentina"))
	fmt.Println(NewTodoServiceIns.AddTodo("Brazil"))
	fmt.Println(NewTodoServiceIns.AddTodo("Romania"))
	//remove
	fmt.Println(NewTodoServiceIns.RemoveTodo("Brazil"))
	//List
	fmt.Println(NewTodoServiceIns.ListTodo())

}
