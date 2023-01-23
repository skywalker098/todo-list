package todo

import (
	"time"
)

// Definig interaface---functionalities
// type todoType interface {
// 	AddTodo(todo string) error
// 	RemoveTodo(todo string) error
// 	ListTodo() []string
// }

// Defining todo "custom data type"
type todo struct {
	Task      string
	Completed bool
	Added     time.Time
}

// todoList store the array of type "todo"
type todoList struct {
	todoStore []todo
}

// returnig tolist{} address
func NewTodoService() *todoList {
	return &todoList{}
}
