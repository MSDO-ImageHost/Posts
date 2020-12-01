package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateOnePost(post PostData) (result PostData, err error) {
	if err := AssertClientInstance(); err != nil {
		return result, err
	}
	return storage.UpdateOne(post)
}

func (s *mongoStorage) UpdateOne(post PostData) (result PostData, err error) {
	// Convert hex string into bson object id
	scaffoldID, err := primitive.ObjectIDFromHex(post.IDHex)
	if err != nil {
		return result, err
	}

	now := time.Now()
	filter := bson.M{"_id": scaffoldID}
	update := bson.D{{Key: "$set", Value: bson.M{"updated_at": now}}}

	// Update header
	if post.Header.Update {
		header := mongoContent{
			ID:        primitive.NewObjectID(),
			Author:    post.Author,
			Data:      post.Header.Data,
			CreatedAt: now,
		}

		_, err = s.HeaderStorage.InsertOne(context.TODO(), header)
		if err != nil {
			return result, err
		}

		update = append(update, bson.E{Key: "$push", Value: bson.M{"headers": header.ID}})
	}

	// Update body
	if post.Body.Update {
		body := mongoContent{
			ID:        primitive.NewObjectID(),
			Author:    post.Author,
			Data:      post.Body.Data,
			CreatedAt: now,
		}

		_, err = s.BodyStorage.InsertOne(context.TODO(), body)
		if err != nil {
			return result, err
		}

		update = append(update, bson.E{Key: "$push", Value: bson.M{"bodies": body.ID}})
	}

	// Update matching document
	_, err = s.ScaffoldStorage.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return result, err
	}

	// Find the updated result
	result, err = s.FindOne(post.IDHex)
	if err != nil {
		return result, err
	}

	return result, nil
}
