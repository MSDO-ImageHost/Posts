package database

import (
	"fmt"

	"github.com/MSDO-ImageHost/Posts/internal/api"
	"github.com/MSDO-ImageHost/Posts/internal/utils"
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
	cur, err := s.ScaffoldStorage.Aggregate(timeOutCtx, aggregation)
	if err != nil {
		fmt.Println(err)
		return results, err
	}
	defer cur.Close(timeOutCtx)

	// Decode findings
	var scaffolds []mongoScaffoldContents
	if err := cur.All(timeOutCtx, &scaffolds); err != nil {
		return results, err
	}

	fmt.Println("Results: ", len(scaffolds))

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
			results[i].Header.AuthorID = scaffolds[0].HeaderContents[0].AuthorID
			results[i].Header.CreatedAt = scaffolds[0].HeaderContents[0].CreatedAt
			results[i].Header.Data = scaffolds[0].HeaderContents[0].Data
		}

		if len(scaffolds[0].BodyContents) > 0 {
			results[i].Body.AuthorID = scaffolds[0].BodyContents[0].AuthorID
			results[i].Body.CreatedAt = scaffolds[0].BodyContents[0].CreatedAt
			results[i].Body.Data = scaffolds[0].BodyContents[0].Data
		}
	}
	fmt.Println(utils.PrettyFormatMap(results))

	defer cancel()
	return results, nil
}

//func (s *mongoStorage) FindMany(postIdHexes []string, paging api.PagingStruct) (results []PostData, err error) {
//
//	// Convert hex strings into bson object ids
//	scaffoldIDs, err := hexes2ObjectID(postIdHexes)
//	if err != nil {
//		return results, err
//	}
//
//	fmt.Println(postIdHexes, scaffoldIDs)
//
//	// Use Mongo aggregation scheme to query documents
//	aggregationScheme := []bson.M{
//		{"$match": bson.M{"_id": bson.M{"$in": scaffoldIDs}}},
//		{"$lookup": bson.M{
//			"from": "headers",
//			"as":   "headers",
//			"let":  bson.D{{Key: "header_refs", Value: "$header_refs"}},
//			"pipeline": mongo.Pipeline{
//				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$header_refs"}}}}},
//				bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}},
//				bson.D{{Key: "$limit", Value: 1}},
//			}},
//		},
//		{"$lookup": bson.M{
//			"from": "bodies",
//			"as":   "bodies",
//			"let":  bson.D{{Key: "body_refs", Value: "$body_refs"}},
//			"pipeline": mongo.Pipeline{
//				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": [2]string{"$_id", "$$body_refs"}}}}},
//				bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}},
//				bson.D{{Key: "$limit", Value: 1}},
//			}},
//		},
//	}
//
//	// Find matching documents
//	cur, err := s.ScaffoldStorage.Aggregate(timeOutCtx, aggregationScheme)
//	if err != nil {
//		return results, err
//	}
//	defer cur.Close(timeOutCtx)
//
//	// Decode findings
//	var scaffolds []mongoScaffoldContents
//	if err := cur.All(timeOutCtx, &scaffolds); err != nil {
//		return results, nil
//	}
//
//	// Update result with data from findings
//	if len(scaffolds) == 0 {
//		return results, fmt.Errorf("No macthing documents found")
//	}
//
//	results = make([]PostData, len(scaffolds))
//	for i := range scaffolds {
//		results[i].IDHex = scaffolds[i].ID.Hex()
//		results[i].AuthorID = scaffolds[i].AuthorID
//		results[i].CreatedAt = scaffolds[i].CreatedAt
//		results[i].UpdatedAt = scaffolds[i].UpdatedAt
//
//		if len(scaffolds[0].HeaderContents) > 0 {
//			results[i].Header.AuthorID = scaffolds[0].HeaderContents[0].AuthorID
//			results[i].Header.CreatedAt = scaffolds[0].HeaderContents[0].CreatedAt
//			results[i].Header.Data = scaffolds[0].HeaderContents[0].Data
//		}
//
//		if len(scaffolds[0].BodyContents) > 0 {
//			results[i].Body.AuthorID = scaffolds[0].BodyContents[0].AuthorID
//			results[i].Body.CreatedAt = scaffolds[0].BodyContents[0].CreatedAt
//			results[i].Body.Data = scaffolds[0].BodyContents[0].Data
//		}
//	}
//
//	fmt.Println("Results: ", len(results))
//
//	return results, nil
//}
//
