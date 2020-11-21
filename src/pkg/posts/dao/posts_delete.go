package dao

import (
	"context"

	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete
func (db *ScaffoldStorage) Delete(postID string) (deletedPostID string, err error) {
	// Fetch and delete the post scaffold
	scaffoldID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return "", err
	}
	var scaffold models.Post
	scaffoldFilter := bson.M{"_id": scaffoldID}
	err = db.ScaffoldCollection.FindOneAndDelete(context.TODO(), scaffoldFilter).Decode(&scaffold)
	if err != nil {
		return "", err
	}

	// Delete all belonging post headers
	headerFilter := bson.M{"_id": bson.M{"$in": scaffold.Header.(primitive.A)}}
	_, err = db.HeaderCollection.DeleteMany(context.TODO(), headerFilter)
	if err != nil {
		return "", err
	}

	// Delete all belonging post bodies
	bodyFilter := bson.M{"_id": bson.M{"$in": scaffold.Body.(primitive.A)}}
	_, err = db.BodyCollection.DeleteMany(context.TODO(), bodyFilter)
	if err != nil {
		return "", err
	}

	return deletedPostID, nil
}
