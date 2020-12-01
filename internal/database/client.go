package database

import (
	"context"
	"fmt"
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
func Init() (err error) {
	log.Println("Database: Setting up")

	// Check if environment variables exists
	log.Println("Database: Checking environment variables")
	if err := CheckEnvs(); err != nil {
		return err
	}
	log.Println("Database: Variables are set")

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

	log.Println("Database: Pinging client")
	if err := Ping(); err != nil {
		return err
	}
	log.Println("Database: Pong from client")

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

func Deinit() (err error) {
	log.Println("Database: Closing connection")

	if err = shell.Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	log.Println("Database: Closed connection")
	return nil
}

func Ping() (err error) {
	if err := shell.Client.Ping(context.TODO(), nil); err != nil {
		return err
	}
	return nil
}

func CheckEnvs() (err error) {

	var envs = []string{"MONGO_CONN_URI", "MONGO_SERVICE_DB"}

	for _, env := range envs {
		if os.Getenv("MONGO_SERVICE_DB") == "" {
			return fmt.Errorf("%s is not configured!", env)
		}
	}
	return nil
}
