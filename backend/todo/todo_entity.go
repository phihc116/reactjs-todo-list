package todo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

type TodoDto struct {
	ID        primitive.ObjectID
	Completed bool
	Body      string
}

type TodoRequestUpdate struct {
	ID        primitive.ObjectID
	Completed bool
	Body      string
}
