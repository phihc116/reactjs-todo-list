package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/phihc116/reactjs-todo-list/db"
	"github.com/phihc116/reactjs-todo-list/todo"
)

func initMongo() {
	conn := os.Getenv("MONGO_CONN")
	dbName := os.Getenv("DB_NAME")

	err := db.InitializeMongoDBContext(conn, dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := fiber.New()

	loadEnv()
	initMongo()

	controller := todo.NewTodoController()
	controller.RegisterRoutes(app)

	PORT := os.Getenv("PORT")
	app.Listen(":" + PORT)
}
