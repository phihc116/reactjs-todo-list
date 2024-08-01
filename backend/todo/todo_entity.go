package todo

type Todo struct {
	ID        int    `json: "id" bson:"_id"`
	Completed bool   `json: "completed"`
	Body      string `json: "body"`
}
