package dao

import (
	"context"
	"fmt"
	"time"

	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Update
func (db *ScaffoldStorage) Update(postID string, post models.Post) (updatedPostID string, err error) {
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

	id, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return "", err
	}

	// Update scaffold
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"header_ids": header.ID, "body_ids": body.ID}}
	_, err = db.ScaffoldCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "", err
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
	return updatedPostID, nil
}
