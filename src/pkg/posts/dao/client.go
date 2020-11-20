package dao

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Init configures the db connection credentials and initializes the database collections
func Init() error {

	// Database handle
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_CONN_STRING")))
	if err != nil {
		return err
	}
	postsDb := client.Database(os.Getenv("MONGO_SERVICE_DB"))

	// Collection handles
	scaffoldColl := postsDb.Collection("scaffolds")
	headerColl := postsDb.Collection("headers")
	bodyColl := postsDb.Collection("bodies")

	// Configure collections and their respective handles
	Posts = &ScaffoldStorage{
		Host:               postsDb,
		ScaffoldCollection: scaffoldColl,
		HeaderCollection:   headerColl,
		BodyCollection:     bodyColl,
	}

	Headers = &HeaderStorage{
		Host:               postsDb,
		ScaffoldCollection: scaffoldColl,
		HeaderCollection:   headerColl,
		BodyCollection:     bodyColl,
	}

	Bodies = &BodyStorage{
		Host:               postsDb,
		ScaffoldCollection: scaffoldColl,
		HeaderCollection:   headerColl,
		BodyCollection:     bodyColl,
	}

	return nil
}

func Deinit() error {
	/*
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()
	*/
	return nil
}
