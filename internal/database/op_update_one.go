package database

import (
	"fmt"
	"time"

	"github.com/MSDO-ImageHost/Posts/internal/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Public module handler
func UpdateOnePost(post PostData, a auth.User) (result PostData, err error) {
	if err := AssertClientInstance(); err != nil {
		return result, err
	}
	return storage.UpdateOne(post, a)
}

func (s *mongoStorage) UpdateOne(post PostData, a auth.User) (result PostData, err error) {
	// Convert hex string into bson object id
	scaffoldID, err := primitive.ObjectIDFromHex(post.IDHex)
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

	now := time.Now()
	filter := bson.M{"_id": scaffoldID}
	update := bson.D{{Key: "$set", Value: bson.M{"updated_at": now}}}

	// Update header
	if post.Header.Update {
		header := mongoContent{
			ID:        primitive.NewObjectID(),
			AuthorID:  a.UserID,
			Data:      post.Header.Data,
			CreatedAt: &now,
		}

		_, err = s.HeaderStorage.InsertOne(timeOutCtx, header)
		if err != nil {
			return result, err
		}

		update = append(update, bson.E{Key: "$push", Value: bson.M{"headers": header.ID}})
	}

	// Update body
	if post.Body.Update {
		body := mongoContent{
			ID:        primitive.NewObjectID(),
			AuthorID:  a.UserID,
			Data:      post.Body.Data,
			CreatedAt: &now,
		}

		_, err = s.BodyStorage.InsertOne(timeOutCtx, body)
		if err != nil {
			return result, err
		}

		update = append(update, bson.E{Key: "$push", Value: bson.M{"bodies": body.ID}})
	}

	// Update matching document
	_, err = s.ScaffoldStorage.UpdateOne(timeOutCtx, filter, update)
	if err != nil {
		return result, err
	}

	// Find the updated result
	result, err = s.FindOne(post.IDHex)
	if err != nil {
		return result, err
	}

	defer cancel()
	return result, nil
}
