package dao

import (
	"context"
	"fmt"
	"time"

	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Adds a new post scaffold
func (db *ScaffoldStorage) Add(post models.Post) (createdPostID string, err error) {
	now := time.Now()

	// Construct post components
	header := models.Content{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Data:      fmt.Sprintf("%v", post.Header),
		CreatedAt: now,
	}

	body := models.Content{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Data:      fmt.Sprintf("%v", post.Body),
		CreatedAt: now,
	}

	scaffold := models.Post{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Header:    []primitive.ObjectID{header.ID},
		Body:      []primitive.ObjectID{body.ID},
		CreatedAt: now,
	}

	// Insert components into their respective collections
	_, err = Headers.Add(header)
	if err != nil {
		return "", err
	}

	_, err = Bodies.Add(body)
	if err != nil {
		return "", err
	}

	res, err := db.ScaffoldCollection.InsertOne(context.TODO(), scaffold)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

// Adds a new post header
func (db *HeaderStorage) Add(content models.Content) (contentID string, err error) {
	res, err := db.HeaderCollection.InsertOne(context.TODO(), content)
	if err != nil {
		return "check error", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

// Adds a new post body
func (db *BodyStorage) Add(content models.Content) (contentID string, err error) {
	res, err := db.BodyCollection.InsertOne(context.TODO(), content)
	if err != nil {
		return "check error", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
