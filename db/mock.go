package db

import (
	"fmt"

	"github.com/amar-jay/surreal/graph/model"
      	"github.com/satori/uuid"
)




type Todos []*model.Todo

func (t *Todos) GetTodos() []*model.Todo {
	return todos
}

// createTodo is the resolver for the createTodo field.
func  (t *Todos) CreateTodo(todo *model.NewTodo) *model.Todo {
  newTodo := &model.Todo{
    ID:   uuid.NewV4().String(),
    Text: todo.Text,
    User: &model.User{
      ID: todo.UserID,
    },
    Done: false,
  };

  todos = append(todos, newTodo)
  return newTodo;
}

// deleteTodo is the resolver for the deleteTodo field.
func  (t *Todos) DeleteTodo(id string) bool {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}

// updateTodo is the resolver for the updateTodo field.
func (t *Todos)  UpdateTodo(id string, todo *model.NewTodo) *model.Todo {
  for i, t := range todos {
      if t.ID == id {
	todos[i].Text = todo.Text
	todos[i].User.ID = todo.UserID
	return todos[i]
      }
  }
    
    return nil
}

// todos is the resolver for the todos field.
func (t *Todos)  ShowTodo() {
	for _, todo := range todos {
		fmt.Println(todo)
	}
}
