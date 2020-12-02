package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Public module handler
func FindOnePost(postIdHex string) (result PostData, err error) {
	if err := AssertClientInstance(); err != nil {
		return result, err
	}
	return storage.FindOne(postIdHex)
}

func (s *mongoStorage) FindOne(postIdHex string) (result PostData, err error) {

	// Convert hex string into bson object id
	scaffoldID, err := primitive.ObjectIDFromHex(postIdHex)
	if err != nil {
		return result, err
	}

	// Use Mongo aggregation scheme to query document
	aggregationScheme := []bson.M{
		{"$match": bson.M{"_id": scaffoldID}},
		//{"$match": bson.M{"$and": []bson.M{{"_id": scaffoldID}, {"marked_deleted": false}}}},
		{"$lookup": bson.M{
			"from": "headers",
			"as":   "headers",
			"let":  bson.D{{Key: "headers", Value: "$headers"}},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$headers"}}}}},
				bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}},
				bson.D{{Key: "$limit", Value: 1}},
			}},
		},
		{"$lookup": bson.M{
			"from": "bodies",
			"as":   "bodies",
			"let":  bson.D{{Key: "bodies", Value: "$bodies"}},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$bodies"}}}}},
				bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}},
				bson.D{{Key: "$limit", Value: 1}},
			}},
		},
	}

	// Find matching document
	cur, err := s.ScaffoldStorage.Aggregate(context.TODO(), aggregationScheme)
	if err != nil {
		return result, err
	}
	defer cur.Close(context.TODO())

	// Decode findings
	var scaffolds []mongoScaffoldContents
	if err := cur.All(context.TODO(), &scaffolds); err != nil {
		return result, nil
	}

	// Update result with data from findings
	if len(scaffolds) == 0 {
		return result, fmt.Errorf("No matching document found")
	}

	result.IDHex = scaffolds[0].ID.Hex()
	result.Author = scaffolds[0].Author
	result.CreatedAt = scaffolds[0].CreatedAt
	result.UpdatedAt = scaffolds[0].UpdatedAt

	if len(scaffolds[0].HeaderContents) > 0 {
		result.Header.Author = scaffolds[0].HeaderContents[0].Author
		result.Header.CreatedAt = scaffolds[0].HeaderContents[0].CreatedAt
		result.Header.Data = scaffolds[0].HeaderContents[0].Data
	}

	if len(scaffolds[0].BodyContents) > 0 {
		result.Body.Author = scaffolds[0].BodyContents[0].Author
		result.Body.CreatedAt = scaffolds[0].BodyContents[0].CreatedAt
		result.Body.Data = scaffolds[0].BodyContents[0].Data
	}
	return result, nil
}
