package todo

import (
	"time"
)

// ADD*************************************************
func (t *todoList) AddTodo(newTodo string) error {
	newtodo := todo{
		Task:      newTodo,
		Completed: false,
		Added:     time.Now(),
	}
	t.todoStore = append(t.todoStore, newtodo)
	//fmt.Println(newtodo)
	return nil
}

// REMOVE*************************************************
func (t *todoList) RemoveTodo(todo string) error {
	for i, v := range t.todoStore {
		if v.Task == todo {
			t.todoStore = append(t.todoStore[:i], t.todoStore[i+1:]...)
		}
	}
	return nil
}

// LIST ALL*************************************************
func (t *todoList) ListTodo() []string {
	var newtodoList []string
	for _, v := range t.todoStore {
		newtodoList = append(newtodoList, v.Task)
	}
	return newtodoList
}
