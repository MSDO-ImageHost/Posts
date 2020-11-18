package dao

import (
	"os"
)

// Setup configures the db connection credentials and initializes the database collections
func Setup() error {
	mongo := MongoDBHost{
		URI:  os.Getenv("MONGO_CONN_STRING"),
		Name: os.Getenv("MONGO_SERVICE_DB"),
	}

	Posts = &ScaffoldStorage{
		Host:       mongo,
		Collection: "posts",
	}

	Titles = &HeaderStorage{
		Host:       mongo,
		Collection: "postheaders",
	}

	Bodies = &BodyStorage{
		Host:       mongo,
		Collection: "postbodies",
	}
	return nil
}
