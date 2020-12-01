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
	log.Println("Database:\tSetting up")

	// Check if environment variables exists
	log.Println("Database:\tChecking environment variables")
	if err := CheckEnvs(); err != nil {
		return err
	}
	log.Println("Database:\tVariables are set")

	// Database handle
	log.Println("Database:\tOpening connection")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_CONN_URI")))
	if err != nil {
		return err
	}
	log.Println("Database:\tConnection opened")

	shell = mongoShell{
		Client: client,
		DB:     client.Database(os.Getenv("MONGO_SERVICE_DB")),
	}

	log.Println("Database:\tPinging client")
	if err := Ping(); err != nil {
		return err
	}
	log.Println("Database:\tPong from client")

	// Configure collections and their respective handles
	log.Println("Database:\tConfiguring collections")
	storage = &mongoStorage{
		ScaffoldStorage: shell.DB.Collection("scaffolds"),
		HeaderStorage:   shell.DB.Collection("headers"),
		BodyStorage:     shell.DB.Collection("bodies"),
		ConsumerStorage: shell.DB.Collection("consumer-meta"),
	}
	log.Println("Database:\tCollections ready")

	log.Println("Database:\tSetup finished")
	return nil
}

func Deinit() (err error) {
	log.Println("Database:\tClosing connection")

	if err = shell.Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	log.Println("Database:\tClosed connection")
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
