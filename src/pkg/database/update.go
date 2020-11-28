package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Update
func (storage *MongoStorage) UpdateOne(post PostScaffold) (updatedPostID string, err error) {
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

	id, err := primitive.ObjectIDFromHex(post.IDHex)
	if err != nil {
		return updatedPostID, err
	}

	// Update scaffold
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"header_ids": header.ID, "body_ids": body.ID}}
	_, err = storage.ScaffoldStorage.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return updatedPostID, err
	}

	// Insert components into their respective collections
	_, err = storage.HeaderStorage.InsertOne(context.TODO(), header)
	if err != nil {
		return updatedPostID, err
	}

	_, err = storage.BodyStorage.InsertOne(context.TODO(), body)
	if err != nil {
		return updatedPostID, err
	}

	return updatedPostID, nil
}
