package dao

import (
	"context"

	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Find
func (db *ScaffoldStorage) Find(postID string) (models.Post, error) {

	// Fetch post scaffold
	scaffoldID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return models.Post{}, err
	}
	var scaffold models.Post
	filter := bson.M{"_id": scaffoldID}
	err = db.ScaffoldCollection.FindOne(context.TODO(), filter).Decode(&scaffold)
	if err != nil {
		return models.Post{}, err
	}

	// Fetch post header
	var header models.Content
	filter = bson.M{"_id": scaffold.Header}
	err = db.HeaderCollection.FindOne(context.TODO(), filter).Decode(&header)
	if err != nil {
		return models.Post{}, err
	}

	// Fetch post body
	var body models.Content
	filter = bson.M{"_id": scaffold.Body}
	err = db.BodyCollection.FindOne(context.TODO(), filter).Decode(&body)
	if err != nil {
		return models.Post{}, err
	}

	fetchedPost := models.Post{
		ID:        scaffold.ID.(primitive.ObjectID).Hex(),
		AuthorID:  scaffold.AuthorID,
		CreatedAt: scaffold.CreatedAt,
		Header:    header.Data,
		Body:      body.Data,
	}

	return fetchedPost, nil

}
