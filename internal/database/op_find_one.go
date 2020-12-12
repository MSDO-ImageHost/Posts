package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	aggregation := make([]bson.M, 0)
	aggregation = append(aggregation, sortByDateDecending)
	aggregation = append(aggregation, headerHistoryDecending(1))
	aggregation = append(aggregation, bodyHistoryDecending(1))
	aggregation = append(aggregation, allOrFilteredScaffoldIds([]primitive.ObjectID{scaffoldID}))

	// Find matching document
	cur, err := s.ScaffoldStorage.Aggregate(timeOutCtx, aggregation)
	if err != nil {
		return result, err
	}
	defer cur.Close(timeOutCtx)

	// Decode findings
	var scaffolds []mongoScaffoldContents
	if err := cur.All(timeOutCtx, &scaffolds); err != nil {
		return result, nil
	}

	// Update result with data from findings
	if len(scaffolds) == 0 {
		return result, fmt.Errorf("No matching document found")
	}

	result.IDHex = scaffolds[0].ID.Hex()
	result.AuthorID = scaffolds[0].AuthorID
	result.CreatedAt = scaffolds[0].CreatedAt
	result.UpdatedAt = scaffolds[0].UpdatedAt

	if len(scaffolds[0].HeaderContents) > 0 {
		result.Header.AuthorID = scaffolds[0].HeaderContents[0].AuthorID
		result.Header.CreatedAt = scaffolds[0].HeaderContents[0].CreatedAt
		result.Header.Data = scaffolds[0].HeaderContents[0].Data
	}

	if len(scaffolds[0].BodyContents) > 0 {
		result.Body.AuthorID = scaffolds[0].BodyContents[0].AuthorID
		result.Body.CreatedAt = scaffolds[0].BodyContents[0].CreatedAt
		result.Body.Data = scaffolds[0].BodyContents[0].Data
	}
	defer cancel()
	return result, nil
}
