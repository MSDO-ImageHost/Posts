package dao

import (
	"context"

	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Find a post in the database with id
func (db *ScaffoldStorage) Find(postID string) (fetchedPost models.Post, err error) {
	// Fetch post scaffold
	scaffoldID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return models.Post{}, err
	}
	var scaffold models.Post
	scaffoldFilter := bson.M{"_id": scaffoldID}
	err = db.ScaffoldCollection.FindOne(context.TODO(), scaffoldFilter).Decode(&scaffold)
	if err != nil {
		return models.Post{}, err
	}

	// Fetch latest post header scaffold.Header.([]primitive.ObjectID)
	var header models.Content
	headerFilter := bson.D{
		{"_id", bson.M{"$in": scaffold.Header.(primitive.A)}},
		//{"$orderby", bson.M{"created_at": -1}}, // TODO: sort by latest
	}
	err = db.HeaderCollection.FindOne(context.TODO(), headerFilter).Decode(&header)
	if err != nil {
		return models.Post{}, err
	}

	// Fetch latest post body
	var body models.Content
	bodyFilter := bson.D{
		{"_id", bson.M{"$in": scaffold.Body.(primitive.A)}},
		//{"$orderby", bson.M{"created_at": -1}}, // TODO: sort by latest
	}

	err = db.BodyCollection.FindOne(context.TODO(), bodyFilter).Decode(&body)
	if err != nil {
		return models.Post{}, err
	}

	// Compose fetched data
	fetchedPost = models.Post{
		ID:        scaffold.ID.(primitive.ObjectID).Hex(),
		AuthorID:  scaffold.AuthorID,
		CreatedAt: scaffold.CreatedAt,
		Header:    header.Data,
		Body:      body.Data,
	}

	return fetchedPost, nil
}
