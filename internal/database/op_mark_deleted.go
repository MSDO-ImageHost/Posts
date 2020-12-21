package database

import (
	"context"

	"github.com/MSDO-ImageHost/Posts/internal/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Public module handler
func MarkDeleteOne(postIdHex string, auth auth.User) (result string, err error) {
	if err := AssertClientInstance(); err != nil {
		return result, err
	}
	return storage.MarkDeleteOne(postIdHex, auth)
}

func (s *mongoStorage) MarkDeleteOne(postIdHex string, auth auth.User) (result string, err error) {
	// Convert hex string into bson object id
	scaffoldID, err := primitive.ObjectIDFromHex(postIdHex)
	if err != nil {
		return result, err
	}

	scaffoldFilter := bson.M{"_id": scaffoldID}
	update := bson.D{{Key: "$set", Value: bson.M{"marked_deleted": true}}}

	// Mark scaffold as deleted
	var scaffoldRef mongoScaffoldRefs
	if err := s.ScaffoldStorage.FindOneAndUpdate(context.TODO(), scaffoldFilter, update).Decode(&scaffoldRef); err != nil {
		return result, err
	}

	// Mark headers as deleted
	headerFilter := bson.M{"_id": bson.M{"$in": scaffoldRef.HeaderRefs}}
	_, err = s.HeaderStorage.UpdateMany(context.TODO(), headerFilter, update)
	if err != nil {
		return result, err
	}

	// Mark bodies as deleted
	bodyFilter := bson.M{"_id": bson.M{"$in": scaffoldRef.BodyRefs}}
	_, err = s.BodyStorage.UpdateMany(context.TODO(), bodyFilter, update)
	if err != nil {
		return result, err
	}

	result = postIdHex
	defer cancel()
	return result, nil
}
