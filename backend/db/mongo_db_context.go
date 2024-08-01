package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCtx *MongoDBContext

type MongoDBContext struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func InitializeMongoDBContext(uri, dbName string) error {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	MongoCtx = &MongoDBContext{
		Client:   client,
		Database: client.Database(dbName),
	}
	return nil
}

func Disconnect() error {
	if MongoCtx != nil {
		return MongoCtx.Client.Disconnect(context.Background())
	}
	return nil
}
