package database

import (
	"fmt"
	"sync"

	"github.com/MSDO-ImageHost/Posts/internal/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Public module handler
func DeleteManyPosts(postIdHexes []string, a auth.User) (results []string, err error) {
	if err := AssertClientInstance(); err != nil {
		return results, err
	}
	return storage.DeleteMany(postIdHexes, a)
}

func (s *mongoStorage) DeleteMany(postIdHexes []string, a auth.User) (results []string, err error) {

	// Convert hex string into bson object id
	scaffoldIDs, err := hexes2ObjectID(postIdHexes)
	if err != nil {
		return results, err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(scaffoldIDs))

	for _, objId := range scaffoldIDs {
		go func(id primitive.ObjectID) error {

			// Permanently delete scaffolds
			var scaffoldRef mongoScaffoldRefs
			if err := s.ScaffoldStorage.FindOne(timeOutCtx, bson.M{"_id": id}).Decode(&scaffoldRef); err != nil {
				return err
			}

			// Check that the user can alter data
			if !a.CanDelete(auth.User{UserID: scaffoldRef.AuthorID}) {
				return fmt.Errorf("Insufficient permissions for id %s", scaffoldRef.ID.Hex())
			}

			// Permanently delete headers and bodies
			s.ScaffoldStorage.DeleteOne(timeOutCtx, bson.M{"_id": id})
			s.HeaderStorage.DeleteMany(timeOutCtx, bson.M{"_id": bson.M{"$in": scaffoldRef.HeaderRefs}})
			s.BodyStorage.DeleteMany(timeOutCtx, bson.M{"_id": bson.M{"$in": scaffoldRef.BodyRefs}})
			results = append(results, scaffoldRef.ID.Hex())

			wg.Done()
			return nil
		}(objId)
	}
	wg.Wait()

	results = postIdHexes
	return results, nil
}
