package todo

import (
	"context"
	"errors"
	"time"

	"github.com/phihc116/reactjs-todo-list/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoInterface interface {
	CreateTodo(todo Todo) (interface{}, error)
	GetList() ([]TodoDto, error)
	UpdateTodo(todoRequestUpdate TodoRequestUpdate) (bool, error)
	DeleteTodo(id primitive.ObjectID) (bool, error)
}

type TodoService struct {
	todos      []Todo
	collection *mongo.Collection
}

func NewTodoService() *TodoService {
	collection := db.MongoCtx.Database.Collection("todos")
	return &TodoService{make([]Todo, 0), collection}
}

func (s *TodoService) CreateTodo(todo Todo) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.collection.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (s *TodoService) GetList() ([]TodoDto, error) {
	var toDos []TodoDto
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := s.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var todo TodoDto
		if err := cur.Decode(&todo); err != nil {
			return nil, err
		}
		toDos = append(toDos, todo)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toDos, nil
}

func (s *TodoService) UpdateTodo(todo TodoRequestUpdate) (bool, error) {
	filter := bson.D{{Key: "_id", Value: todo.ID}}

	update := bson.D{{
		Key: "$set", Value: bson.D{{Key: "completed", Value: todo.Completed}, {Key: "body", Value: todo.Body}},
	}}

	result, err := s.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, errors.New("no documents matched the provided ID")
	}

	return true, nil
}

func (s *TodoService) DeleteTodo(id primitive.ObjectID) (bool, error) {

	filter := bson.D{{Key: "_id", Value: id}}

	_, err := s.collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return false, err
	}

	return true, nil
}
