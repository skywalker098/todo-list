package todo

import (
	"fmt"
	"time"
)

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

// ADD*************************************************
func (t *todoList) AddTodo(str string) error {
	t.LoadFromJson()
	// fmt.Println("data:", t.todoStore)

	newtodo := todo{
		Task:      str,
		Completed: false,
		Added:     time.Now(),
	}
	t.todoStore = append(t.todoStore, newtodo)
	//fmt.Println(newtodo)
	t.SavetoJson()
	return nil
}

// REMOVE*************************************************
func (t *todoList) RemoveTodo(todo string) error {
	t.LoadFromJson()
	for i, v := range t.todoStore {
		if v.Task == todo {
			t.todoStore = append(t.todoStore[:i], t.todoStore[i+1:]...)
		}
	}
	fmt.Println(t.todoStore)
	t.SavetoJson()
	return nil

}

// LIST ALL*************************************************
func (t *todoList) ListTodo() {
	t.LoadFromJson()
	var newtodoList []string
	for _, v := range t.todoStore {
		newtodoList = append(newtodoList, v.Task)
	}
	fmt.Println(newtodoList)
}
