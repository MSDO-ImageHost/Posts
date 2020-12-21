package database

import (
	"context"
	"fmt"

	"github.com/MSDO-ImageHost/Posts/internal/api"
	"go.mongodb.org/mongo-driver/bson"
)

// Public module handler
func FindManyPosts(postIdHexes []string, paging api.PagingStruct) (results []PostData, err error) {
	if err := AssertClientInstance(); err != nil {
		return results, err
	}
	return storage.FindMany(postIdHexes, paging)
}

func (s *mongoStorage) FindMany(postIdHexes []string, paging api.PagingStruct) (results []PostData, err error) {

	// Convert hex strings into bson object ids
	scaffoldIDs, err := hexes2ObjectID(postIdHexes)
	if err != nil {
		return results, err
	}

	// Use Mongo aggregation scheme to query documents
	aggregation := make([]bson.M, 0)
	aggregation = append(aggregation, sortByDateDecending)
	aggregation = append(aggregation, headerHistoryDecending(1))
	aggregation = append(aggregation, bodyHistoryDecending(1))
	aggregation = append(aggregation, allOrFilteredScaffoldIds(scaffoldIDs))
	//aggregation = append(aggregation, paginationBuilder(paging)...)

	// Find matching documents
	cur, err := s.ScaffoldStorage.Aggregate(context.TODO(), aggregation)
	if err != nil {
		fmt.Println(err)
		return results, err
	}
	defer cur.Close(context.TODO())

	// Decode findings
	var scaffolds []mongoScaffoldContents
	if err := cur.All(context.TODO(), &scaffolds); err != nil {
		return results, err
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

	defer cancel()
	return results, nil
}
