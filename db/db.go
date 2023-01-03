package db

import (
	"context"
	"fmt"

	"github.com/amar-jay/surreal/graph/model"
	surreal "github.com/garrison-henkle/surrealdb.go"
	"github.com/satori/uuid"
)



var (
  todos = []*model.Todo{};
)

type Repository interface {
  GetTodos() []*model.Todo
  CreateTodo(todo *model.NewTodo) *model.Todo
  DeleteTodo(id string) bool
  UpdateTodo(id string, todo *model.NewTodo) *model.Todo
  ShowTodo()
}

type TodoRepo struct {
  Db *surreal.DB
  Ctx context.Context
}

func Init(db surreal.DB) *TodoRepo {

  return &TodoRepo{
		Db: &db,
	}
}

func (t *TodoRepo) GetTodos() []*model.Todo {
  res := t.Db.Select(&t.Ctx, "todo")
  list := []*model.Todo{}
  res.Unmarshal(&list)

  return list
}

// createTodo is the resolver for the createTodo field.
func (t *TodoRepo) CreateTodo(todo *model.NewTodo) (*model.Todo, error) {
  newTodo := &model.Todo{
    ID:   uuid.NewV4().String(),
    Text: todo.Text,
    User: &model.User{
      ID: todo.UserID,
    },
    Done: false,
  };

  res := t.Db.Create(&t.Ctx, "todo", &newTodo)
  if res.Error != nil {
    return nil, res.Error
    }
  todos = append(todos, newTodo)
  return newTodo, nil;
}

// deleteTodo is the resolver for the deleteTodo field.
func (t *TodoRepo) DeleteTodo(id string) (bool, error) {
  res := t.Db.Delete(&t.Ctx, id)
  if res != nil { return false, res
  }
	return true, nil;
}

// updateTodo is the resolver for the updateTodo field.
func (t *TodoRepo) UpdateTodo(id string, todo *model.NewTodo) *model.Todo {
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
func (t *TodoRepo) ShowTodo() {
	for _, todo := range todos {
		fmt.Println(todo)
	}
}
