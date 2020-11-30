package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (storage *MongoStorage) Add(post PostScaffold) (string, error) {

	now := time.Now()

	// Construct post components
	header := Content{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Data:      post.HeaderContent,
		CreatedAt: now,
	}

	body := Content{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Data:      post.BodyContent,
		CreatedAt: now,
	}

	// Update post object with reference and missing data
	post.ID = primitive.NewObjectID()
	post.Header = []primitive.ObjectID{header.ID}
	post.Body = []primitive.ObjectID{body.ID}
	post.CreatedAt = now

	// Insert components into their respective collections
	_, err := storage.HeaderStorage.InsertOne(context.TODO(), header)
	if err != nil {
		return "", err
	}

	_, err = storage.BodyStorage.InsertOne(context.TODO(), body)
	if err != nil {
		return "", err
	}

	res, err := storage.ScaffoldStorage.InsertOne(context.TODO(), post)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
