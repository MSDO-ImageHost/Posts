package database

import (
	"fmt"

	"github.com/MSDO-ImageHost/Posts/internal/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	sortByDateDecending = bson.M{"$sort": bson.M{"created_at": -1}}
)

func headerHistoryDecending(limit uint) bson.M {
	return bson.M{
		"$lookup": bson.M{
			"from": "headers",
			"as":   "headers",
			"let":  bson.D{{Key: "headers", Value: "$headers"}},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$headers"}}}}},
				bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}},
				bson.D{{Key: "$limit", Value: limit}},
			},
		},
	}
}

func bodyHistoryDecending(limit uint) bson.M {
	return bson.M{
		"$lookup": bson.M{
			"from": "bodies",
			"as":   "bodies",
			"let":  bson.D{{Key: "bodies", Value: "$bodies"}},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$bodies"}}}}},
				bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}},
				bson.D{{Key: "$limit", Value: limit}},
			},
		},
	}
}

// Builds an aggregation filter that matches all documents or only those specified
func allOrFilteredScaffoldIds(objIds []primitive.ObjectID) bson.M {
	if len(objIds) < 1 {
		return bson.M{"$match": bson.M{"_id": bson.M{"$exists": true}}}
	}
	return bson.M{"$match": bson.M{"_id": bson.M{"$in": objIds}}}
}

// Builds aggregation filter based on specified pagination
func paginationBuilder(p api.PagingStruct) []bson.M {

	pagination := make([]bson.M, 0)

	// Decending sort by date
	pagination = append(pagination, bson.M{"$sort": bson.M{"created_at": -1}})

	if p.TimeBased() {
		startTime, err := p.ParseStartTime()
		if err == nil {
			filterGreaterThan := bson.M{"$project": bson.M{"created_at": bson.M{"$gt": startTime}}}
			pagination = append(pagination, filterGreaterThan)
			fmt.Println("wqeweqw", filterGreaterThan)
		}
	}

	pagination = append(pagination, bson.M{"$limit": p.Limit})
	return pagination
}
