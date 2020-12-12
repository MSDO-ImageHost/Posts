package database

import (
	"fmt"

	"github.com/MSDO-ImageHost/Posts/internal/api"
	"go.mongodb.org/mongo-driver/bson"
)

// Public module handler
func FindUserPosts(author_id string, paging api.PagingStruct) (results []PostData, err error) {
	if err := AssertClientInstance(); err != nil {
		return results, err
	}
	return storage.FindUserPosts(author_id, paging)
}

func (s *mongoStorage) FindUserPosts(author_id string, paging api.PagingStruct) (results []PostData, err error) {

	// Use Mongo aggregation scheme to query document
	aggregation := make([]bson.M, 0)
	aggregation = append(aggregation, sortByDateDecending)
	aggregation = append(aggregation, headerHistoryDecending(1))
	aggregation = append(aggregation, bodyHistoryDecending(1))
	aggregation = append(aggregation, bson.M{"$match": bson.M{"creator_id": author_id}})

	// Find matching documents
	cur, err := s.ScaffoldStorage.Aggregate(timeOutCtx, aggregation)
	if err != nil {
		return results, err
	}
	defer cur.Close(timeOutCtx)

	// Decode findings
	var scaffolds []mongoScaffoldContents
	if err := cur.All(timeOutCtx, &scaffolds); err != nil {
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
	defer cancel()
	return results, nil
}
