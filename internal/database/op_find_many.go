package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Public module handler
func FindManyPosts(postIdHexes []string) (results []PostData, err error) {
	if err := AssertClientInstance(); err != nil {
		return results, err
	}
	return storage.FindMany(postIdHexes)

}

func (s *mongoStorage) FindMany(postIdHexes []string) (results []PostData, err error) {

	// Convert hex strings into bson object ids
	scaffoldIDs, err := hexes2ObjectID(postIdHexes)
	if err != nil {
		return results, err
	}

	// Use Mongo aggregation scheme to query documents
	aggregationScheme := []bson.M{
		{"$match": bson.M{"_id": bson.M{"$in": scaffoldIDs}}},
		{"$lookup": bson.M{
			"from": "headers",
			"as":   "headers",
			"let":  bson.D{{Key: "header_refs", Value: "$header_refs"}},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$header_refs"}}}}},
				bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}},
				bson.D{{Key: "$limit", Value: 1}},
			}},
		},
		{"$lookup": bson.M{
			"from": "bodies",
			"as":   "bodies",
			"let":  bson.D{{Key: "body_refs", Value: "$body_refs"}},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$body_refs"}}}}},
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
