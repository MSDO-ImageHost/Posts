package database

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Public module handler
func DeleteManyPosts(postIdHexes []string) (results []string, err error) {
	if err := AssertClientInstance(); err != nil {
		return results, err
	}
	return storage.DeleteMany(postIdHexes)
}

func (s *mongoStorage) DeleteMany(postIdHexes []string) (results []string, err error) {

	// Convert hex string into bson object id
	scaffoldIDs, err := hexes2ObjectID(postIdHexes)
	if err != nil {
		return results, err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(scaffoldIDs))

	var errors error
	for _, objId := range scaffoldIDs {
		go func(id primitive.ObjectID) {

			// Permanently delete scaffolds
			var scaffoldRef mongoScaffoldRefs
			if err := s.ScaffoldStorage.FindOneAndDelete(context.TODO(), bson.M{"_id": id}).Decode(&scaffoldRef); err != nil {
				errors = err
			}

			// Permanently delete headers
			_, err = s.HeaderStorage.DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": scaffoldRef.HeaderRefs}})
			if err != nil {
				errors = err
			}

			// Permanently delete bodies
			_, err = s.BodyStorage.DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": scaffoldRef.BodyRefs}})
			if err != nil {
				errors = err
			}
			wg.Done()
		}(objId)
	}
	wg.Wait()

	results = postIdHexes
	return results, errors
}
