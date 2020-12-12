package database

import (
	"fmt"

	auth "github.com/MSDO-ImageHost/Posts/internal/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Public module handler
func DeleteOnePost(postIdHex string, a auth.User) (result PostData, err error) {
	if err := AssertClientInstance(); err != nil {
		return result, err
	}
	return storage.DeleteOne(postIdHex, a)
}

func (s *mongoStorage) DeleteOne(postIdHex string, a auth.User) (result PostData, err error) {

	// Convert hex string into bson object id
	scaffoldID, err := primitive.ObjectIDFromHex(postIdHex)
	if err != nil {
		return result, err
	}

	// Find the scaffold
	var scaffoldRef mongoScaffoldRefs
	if err := s.ScaffoldStorage.FindOne(timeOutCtx, bson.M{"_id": scaffoldID}).Decode(&scaffoldRef); err != nil {
		return result, err
	}

	// Check that the user can alter data
	if canModify := a.CanModify(auth.User{UserID: scaffoldRef.AuthorID}); !canModify {
		return result, fmt.Errorf("Insufficient permissions")
	}

	result, err = s.FindOne(postIdHex)
	if err != nil {
		return result, err
	}

	// Permanently delete scaffold
	if err := s.ScaffoldStorage.FindOneAndDelete(timeOutCtx, bson.M{"_id": scaffoldRef.ID}).Decode(&scaffoldRef); err != nil {
		return result, err
	}

	// Permanently delete headers
	_, err = s.HeaderStorage.DeleteMany(timeOutCtx, bson.M{"_id": bson.M{"$in": scaffoldRef.HeaderRefs}})
	if err != nil {
		return result, err
	}

	// Permanently delete bodies
	_, err = s.BodyStorage.DeleteMany(timeOutCtx, bson.M{"_id": bson.M{"$in": scaffoldRef.BodyRefs}})
	if err != nil {
		return result, err
	}

	defer cancel()
	return result, nil
}
