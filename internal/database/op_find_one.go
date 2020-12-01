package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOnePost(postIdHex string) (result PostData, err error) {
	result, err = storage.FindOne(postIdHex)
	if err != nil {
		// TODO: catch timeout errors
	}
	return result, err
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
		{"$lookup": bson.M{
			"from": "headers",
			"as":   "headers",
			"let":  bson.D{{Key: "header_ids", Value: "$header_ids"}},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$header_ids"}}}}},
				bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}},
				bson.D{{Key: "$limit", Value: 1}},
			}},
		},
		{"$lookup": bson.M{
			"from": "bodies",
			"as":   "bodies",
			"let":  bson.D{{Key: "body_ids", Value: "$body_ids"}},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$body_ids"}}}}},
				bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}},
				bson.D{{Key: "$limit", Value: 1}},
			}},
		},
	}

	// Find matching documents
	cur, err := s.ScaffoldStorage.Aggregate(context.TODO(), aggregationScheme)
	if err != nil {
		return result, err
	}
	defer cur.Close(context.TODO())

	// Decode findings
	var scaffolds []mongoScaffold
	if err := cur.All(context.TODO(), &scaffolds); err != nil {
		return result, nil
	}

	// Update result with data from findings
	if len(scaffolds) == 0 {
		return result, fmt.Errorf("No document with ID %s found", postIdHex)
	}

	result.IDHex = scaffolds[0].ID.Hex()
	result.Author = scaffolds[0].Author
	result.CreatedAt = scaffolds[0].CreatedAt
	result.UpdatedAt = scaffolds[0].UpdatedAt

	if len(scaffolds[0].Headers) > 0 {
		result.Header.Author = scaffolds[0].Headers[0].Author
		result.Header.CreatedAt = scaffolds[0].Headers[0].CreatedAt
		result.Header.Data = scaffolds[0].Headers[0].Data
	}

	if len(scaffolds[0].Bodies) > 0 {
		result.Body.Author = scaffolds[0].Bodies[0].Author
		result.Body.CreatedAt = scaffolds[0].Bodies[0].CreatedAt
		result.Body.Data = scaffolds[0].Bodies[0].Data
	}
	return result, nil
}
