package dao

import (
	"context"
	"time"

	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Adds a new post scaffold
func (db *ScaffoldStorage) Add(post models.Post) (string, error) {

	loc, _ := time.LoadLocation(post.Meta.TZ)
	now := time.Now().In(loc)

	// Construct post components
	header := Content{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Data:      post.Header,
		CreatedAt: now,
	}

	body := Content{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Data:      post.Body,
		CreatedAt: now,
	}

	scaffold := Scaffold{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Header:    header.ID,
		Body:      body.ID,
		CreatedAt: now,
	}

	// Insert components into their respective collections
	_, err := Headers.Add(header)
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
func (db *HeaderStorage) Add(content Content) (string, error) {

	res, err := db.HeaderCollection.InsertOne(context.TODO(), content)
	if err != nil {
		return "check error", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

// Adds a new post body
func (db *BodyStorage) Add(content Content) (string, error) {

	res, err := db.BodyCollection.InsertOne(context.TODO(), content)
	if err != nil {
		return "check error", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
