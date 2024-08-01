package todo

import (
	"errors"

	"github.com/phihc116/reactjs-todo-list/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoInterface interface {
	CreateTodo(todo Todo) error
	GetList() []Todo
	UpdateTodo(id int) (Todo, error)
	DeleteTodo(id int) error
}

type TodoService struct {
	todos      []Todo
	collection *mongo.Collection
}

func NewTodoService() *TodoService {
	collection := db.MongoCtx.Database.Collection("todos")
	return &TodoService{make([]Todo, 0), collection}
}

func (s *TodoService) CreateTodo(todo Todo) error {
	if todo.Body == "" {
		return errors.New("todo body cannot be empty")
	}
	todo.ID = len(s.todos) + 1
	s.todos = append(s.todos, todo)
	return nil
}

func (s *TodoService) GetList() []Todo {
	return s.todos
}

func (s *TodoService) UpdateTodo(id int) (Todo, error) {
	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos[i].Completed = !s.todos[i].Completed
			return s.todos[i], nil
		}
	}
	return Todo{}, errors.New("todo not found")
}

func (s *TodoService) DeleteTodo(id int) error {
	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
