package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	shell MongoShell
	Posts Interface
)

// Init configures the db connection credentials and initializes the database collections.
func Init() error {
	log.Println("Opening database connection")

	// Database handle
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_CONN_URI")))
	if err != nil {
		return err
	}

	shell = MongoShell{
		Client: client,
		DB:     client.Database(os.Getenv("MONGO_SERVICE_DB")),
	}

	// Configure collections and their respective handles
	Posts = &MongoStorage{
		ScaffoldStorage: shell.DB.Collection("scaffolds"),
		HeaderStorage:   shell.DB.Collection("headers"),
		BodyStorage:     shell.DB.Collection("bodies"),
	}

	log.Println("Opened database connection")

	return nil
}

func Deinit() error {
	log.Println("Closing database connection")

	if err := shell.Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	log.Println("Closed database connection")
	return nil
}
