package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete
func (storage *MongoStorage) DeleteOne(postID string) (deletedID string, err error) {

	// Fetch and delete the post scaffold
	scaffoldID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return deletedID, err
	}
	var scaffold PostScaffold
	scaffoldFilter := bson.M{"_id": scaffoldID}
	err = storage.ScaffoldStorage.FindOneAndDelete(context.TODO(), scaffoldFilter).Decode(&scaffold)
	if err != nil {
		return deletedID, err
	}

	// Delete all belonging post headers
	headerFilter := bson.M{"_id": bson.M{"$in": scaffold.Header}}
	_, err = storage.HeaderStorage.DeleteMany(context.TODO(), headerFilter)
	if err != nil {
		return deletedID, err
	}

	// Delete all belonging post bodies
	bodyFilter := bson.M{"_id": bson.M{"$in": scaffold.Body}}
	_, err = storage.BodyStorage.DeleteMany(context.TODO(), bodyFilter)
	if err != nil {
		return deletedID, err
	}

	return deletedID, nil
}
