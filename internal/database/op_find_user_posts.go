package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Public module handler
func FindUserPosts(author string) (results []PostData, err error) {
	if err := AssertClientInstance(); err != nil {
		return results, err
	}
	return storage.FindUserPosts(author)
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
		results[i].AuthorID = scaffolds[i].AuthorID
		results[i].CreatedAt = scaffolds[i].CreatedAt
		results[i].UpdatedAt = scaffolds[i].UpdatedAt

		if len(scaffolds[0].HeaderContents) > 0 {
			results[i].Header.AuthorID = scaffolds[i].HeaderContents[0].AuthorID
			results[i].Header.CreatedAt = scaffolds[i].HeaderContents[0].CreatedAt
			results[i].Header.Data = scaffolds[i].HeaderContents[0].Data
		}

		if len(scaffolds[0].BodyContents) > 0 {
			results[i].Body.AuthorID = scaffolds[i].BodyContents[0].AuthorID
			results[i].Body.CreatedAt = scaffolds[i].BodyContents[0].CreatedAt
			results[i].Body.Data = scaffolds[i].BodyContents[0].Data
		}
	}

	return results, nil
}
