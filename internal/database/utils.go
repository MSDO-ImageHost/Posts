package database

import (
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidHexFormat(hexId string) (err error) {
	_, err = primitive.ObjectIDFromHex(hexId)
	return err
}

func hexes2ObjectID(hexIds []string) (objIds []primitive.ObjectID, err error) {

	// New slice for Object Ids
	objIds = make([]primitive.ObjectID, len(hexIds))

	// Map hex id to object id
	for i, hexId := range hexIds {
		objId, err := primitive.ObjectIDFromHex(hexId)
		if err != nil {
			return nil, err
		}
		objIds[i] = objId
	}
	return objIds, nil
}

func PrettyFormatMap(d interface{}) string {
	b, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

/** landfill

aggregationScheme := []bson.M{
	{"$match": bson.M{"_id": scaffoldID}},
	{"$set": bson.M{"updated_at": now}},
	{"$set": bson.M{"headers": bson.M{"$concatArrays": bson.A{"$headers", []primitive.ObjectID{header.ID}}}}},
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

**/
