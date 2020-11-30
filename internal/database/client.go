package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	shell   mongoShell
	storage storageInterface
)

// Init configures the db connection credentials and initializes the database collections.
func Init() error {
	log.Println("Database: Setting up")

	// Database handle
	log.Println("Database: Opening connection")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_CONN_URI")))
	if err != nil {
		return err
	}
	log.Println("Database: Connection opened")

	shell = mongoShell{
		Client: client,
		DB:     client.Database(os.Getenv("MONGO_SERVICE_DB")),
	}

	// Configure collections and their respective handles
	log.Println("Database: Configuring collections")
	storage = &mongoStorage{
		ScaffoldStorage: shell.DB.Collection("scaffolds"),
		HeaderStorage:   shell.DB.Collection("headers"),
		BodyStorage:     shell.DB.Collection("bodies"),
		ConsumerStorage: shell.DB.Collection("consumer-meta"),
	}
	log.Println("Database: Collections ready")

	log.Println("Database: Setup finished")
	return nil
}

func Deinit() error {
	log.Println("Database: Closing connection")

	if err := shell.Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	log.Println("Database: Closed connection")
	return nil
}
