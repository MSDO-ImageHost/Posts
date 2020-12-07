package database

import (
	"context"
	"fmt"
	"log"

	"github.com/MSDO-ImageHost/Posts/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (pd *PostData) Find() error {

	// Convert hex string into bson object id
	scaffoldID, err := primitive.ObjectIDFromHex(pd.ID.(string))
	if err != nil {
		return err
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
	cur, err := db.ScaffoldStorage.Aggregate(context.TODO(), aggregationScheme)
	if err != nil {
		return err
	}
	defer cur.Close(context.TODO())

	// Decode findings
	var scaffolds []PostData
	if err := cur.All(context.TODO(), &scaffolds); err != nil {
		return nil
	}

	// Update result with data from findings
	if len(scaffolds) == 0 {
		return fmt.Errorf("No matching document found")
	}

	log.Println(_LOG_TAG, utils.PrettyFormatMap(scaffolds))

	// Populate post data with found info
	var scaffold PostData = scaffolds[0]
	header, err := ParseMongoPostContent(scaffold.Headers)
	if err != nil {
		return err
	}
	body, err := ParseMongoPostContent(scaffold.Bodies)
	if err != nil {
		return err
	}

	pd.ID = scaffold.ID.(primitive.ObjectID).Hex()
	pd.AuthorID = scaffold.AuthorID
	pd.CreatedAt = scaffold.CreatedAt
	pd.UpdatedAt = scaffold.UpdatedAt
	pd.Headers = header
	pd.Bodies = body
	return nil
}
