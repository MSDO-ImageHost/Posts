package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserPosts(author string) (results []PostData, err error) {
	results, err = storage.FindUserPosts(author)
	if err != nil {
		// TODO: catch timeout errors
	}
	return results, err
}

func (s *mongoStorage) FindUserPosts(author string) (results []PostData, err error) {

	// Use Mongo aggregation scheme to query documents
	aggregationScheme := []bson.M{
		{"$match": bson.M{"creator_id": author}},
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

	// Find matching documents
	cur, err := s.ScaffoldStorage.Aggregate(context.TODO(), aggregationScheme)
	if err != nil {
		return results, err
	}
	defer cur.Close(context.TODO())

	// Decode findings
	var scaffolds []mongoScaffoldContents
	if err := cur.All(context.TODO(), &scaffolds); err != nil {
		return results, nil
	}

	// Update result with data from findings
	if len(scaffolds) == 0 {
		return results, fmt.Errorf("No macthing documents found")
	}

	results = make([]PostData, len(scaffolds))
	for i := range scaffolds {
		results[i].IDHex = scaffolds[i].ID.Hex()
		results[i].Author = scaffolds[i].Author
		results[i].CreatedAt = scaffolds[i].CreatedAt
		results[i].UpdatedAt = scaffolds[i].UpdatedAt

		if len(scaffolds[0].HeaderContents) > 0 {
			results[i].Header.Author = scaffolds[0].HeaderContents[0].Author
			results[i].Header.CreatedAt = scaffolds[0].HeaderContents[0].CreatedAt
			results[i].Header.Data = scaffolds[0].HeaderContents[0].Data
		}

		if len(scaffolds[0].BodyContents) > 0 {
			results[i].Body.Author = scaffolds[0].BodyContents[0].Author
			results[i].Body.CreatedAt = scaffolds[0].BodyContents[0].CreatedAt
			results[i].Body.Data = scaffolds[0].BodyContents[0].Data
		}
	}

	return results, nil
}
