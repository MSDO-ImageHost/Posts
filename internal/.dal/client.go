package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db Database

// Connect initializes a new database connection session
func Connect(dbUri string) error {
	// Dial Mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		return err
	}
	database := client.Database(DatabaseName)

	// Mongo Shell object
	db = Database{
		Client:          client,
		Db:              database,
		ScaffoldStorage: database.Collection(ScaffoldCollectionName),
		HeaderStorage:   database.Collection(HeaderCollectionName),
		BodyStorage:     database.Collection(BodyCollectionName),
	}
	return nil
}

// Disconnect from Mongo instance
func Disconnect() (err error) {
	if err = db.Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}

// Test if Mongo instance and connection is still alive
func Ping() (err error) {
	if err := db.Client.Ping(context.TODO(), nil); err != nil {
		return err
	}
	return nil
}
