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
func Init() (err error) {
	log.Println(_LOG_TAG, "Setting up")

	// Database handle
	log.Println(_LOG_TAG, "Opening connection")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_CONN_URI")))
	//client, err := mongo.Connect(context.TODO(), &options.ClientOptions{
	//	Hosts:
	//	Auth: &options.Credential{
	//		Username:    "rooter",
	//		Password:    "toorer",
	//		PasswordSet: true,
	//	},
	//})

	if err != nil {
		return err
	}
	log.Println(_LOG_TAG, "Connection opened")

	shell = mongoShell{
		Client: client,
		DB:     client.Database("posts"),
	}

	log.Println(_LOG_TAG, "Pinging server")
	if err := Ping(); err != nil {
		return err
	}
	log.Println(_LOG_TAG, "Pong from server")

	// Configure collections and their respective handles
	log.Println(_LOG_TAG, "Configuring collections")
	storage = &mongoStorage{
		ScaffoldStorage: shell.DB.Collection("scaffolds"),
		HeaderStorage:   shell.DB.Collection("headers"),
		BodyStorage:     shell.DB.Collection("bodies"),
		ConsumerStorage: shell.DB.Collection("consumer-meta"),
	}
	log.Println(_LOG_TAG, "Collections ready")

	log.Println(_LOG_TAG, "Setup finished")
	defer cancel()
	return nil
}

func Deinit() (err error) {
	log.Println(_LOG_TAG, "Closing connection")

	if err = shell.Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	log.Println(_LOG_TAG, "Closed connection")
	defer cancel()
	return nil
}

func Ping() (err error) {
	if err := shell.Client.Ping(context.TODO(), nil); err != nil {
		return err
	}
	defer cancel()
	return nil
}
