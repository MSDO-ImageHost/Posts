package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Public module handler
func FindHistoryPost(postIdHex string) (result PostDataHistory, err error) {
	if err := AssertClientInstance(); err != nil {
		return result, err
	}
	return storage.FindHistory(postIdHex)
}

func (s *mongoStorage) FindHistory(postIdHex string) (result PostDataHistory, err error) {

	// Convert hex string into bson object id
	scaffoldID, err := primitive.ObjectIDFromHex(postIdHex)
	if err != nil {
		return result, err
	}

	// Use Mongo aggregation scheme to query document
	aggregationScheme := []bson.M{
		{"$match": bson.M{"_id": scaffoldID}},
		{"$lookup": bson.M{
			"from": "headers",
			"as":   "headers",
			"let":  bson.D{{Key: "headers", Value: "$headers"}},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$headers"}}}}},
				bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}}}},
		},
		{"$lookup": bson.M{
			"from": "bodies",
			"as":   "bodies",
			"let":  bson.D{{Key: "bodies", Value: "$bodies"}},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$bodies"}}}}},
				bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}},
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
		headers := make([]PostContent, len(scaffolds[0].HeaderContents))
		for i, content := range scaffolds[0].HeaderContents {
			headers[i].Author = content.Author
			headers[i].CreatedAt = content.CreatedAt
			headers[i].Data = content.Data
		}
		result.Headers = headers
	}

	if len(scaffolds[0].BodyContents) > 0 {
		bodies := make([]PostContent, len(scaffolds[0].BodyContents))
		for i, content := range scaffolds[0].BodyContents {
			bodies[i].Author = content.Author
			bodies[i].CreatedAt = content.CreatedAt
			bodies[i].Data = content.Data
		}
		result.Bodies = bodies
	}

	return result, nil
}
