package main

import (
	"github.com/todolist_iteration1/todo"
)

func main() {
	NewTodoServiceIns := todo.NewTodoService()
	// NewTodoServiceIns.LoadFromJson()

	//add
	// fmt.Println(NewTodoServiceIns.AddTodo("Argentina"))
	// fmt.Println(NewTodoServiceIns.AddTodo("Brazil"))
	// fmt.Println(NewTodoServiceIns.AddTodo("Romania"))
	// fmt.Println(NewTodoServiceIns.AddTodo("São Paulo"))
	// NewTodoServiceIns.AddTodo("France")
	NewTodoServiceIns.ListTodo()

	//remove
	// NewTodoServiceIns.RemoveTodo("Brazil")
	//List
	// NewTodoServiceIns.ListTodo()

}
