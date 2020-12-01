package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteOnePost(postIdHex string) (result string, err error) {
	result, err = storage.DeleteOne(postIdHex)
	if err != nil {
		// TODO: catch timeout errors
	}
	return result, err
}

func (s *mongoStorage) DeleteOne(postIdHex string) (result string, err error) {

	// Convert hex string into bson object id
	scaffoldID, err := primitive.ObjectIDFromHex(postIdHex)
	if err != nil {
		return result, err
	}

	// Permanently delete scaffolds
	var scaffoldRef mongoScaffoldRefs
	if err := s.ScaffoldStorage.FindOneAndDelete(context.TODO(), bson.M{"_id": scaffoldID}).Decode(&scaffoldRef); err != nil {
		return result, err
	}

	// Permanently delete headers
	_, err = s.HeaderStorage.DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": scaffoldRef.HeaderRefs}})
	if err != nil {
		return result, err
	}

	// Permanently delete bodies
	_, err = s.BodyStorage.DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": scaffoldRef.BodyRefs}})
	if err != nil {
		return result, err
	}

	result = postIdHex
	return result, nil
}
