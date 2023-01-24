package todo

import (
	"fmt"
	"time"
)

// Defining todo "custom data type"
type todo struct {
	Id        int
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

// -------------Add Todo --------------------------------
func (t *todoList) AddTodo(str string) error {
	t.LoadFromJson()
	// fmt.Println("data:", t.todoStore)

	newtodo := todo{
		Id:        t.AddnewId(),
		Task:      str,
		Completed: false,
		Added:     time.Now(),
	}
	t.todoStore = append(t.todoStore, newtodo)
	//fmt.Println(newtodo)
	fmt.Printf("data added: %v\n", str)

	t.SavetoJson()
	return nil
}

// // -------------Remove Todo --------------------------------
func (t *todoList) RemoveTodo(id int) error {
	t.LoadFromJson()
	for i, v := range t.todoStore {
		if v.Id == id {
			t.todoStore = append(t.todoStore[:i], t.todoStore[i+1:]...)
		}
	}
	fmt.Println("data removed with id:", id)
	t.SavetoJson()
	return nil

}

// -------------List Todo ------------------------------------
func (t *todoList) ListTodo() {
	t.LoadFromJson()
	var newtodoList []string
	for _, v := range t.todoStore {
		newtodoList = append(newtodoList, v.Task)
	}
	fmt.Println(newtodoList)
}
