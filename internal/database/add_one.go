package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddOnePost(post PostData) (result PostData, err error) {
	result, err = storage.AddOne(post)
	if err != nil {
		// TODO: catch timeout errors
	}

	return result, err
}

func (s *mongoStorage) AddOne(post PostData) (result PostData, err error) {

	now := time.Now()

	// Construct post content components
	header := mongoContent{
		ID:        primitive.NewObjectID(),
		OwnedBy:   post.Author,
		Data:      post.Header,
		CreatedAt: now,
	}

	body := mongoContent{
		ID:        primitive.NewObjectID(),
		OwnedBy:   post.Author,
		Data:      post.Body,
		CreatedAt: now,
	}

	scaffold := mongoScaffold{
		ID:         primitive.NewObjectID(),
		OwnedBy:    post.Author,
		CreatedAt:  now,
		HeaderRefs: []primitive.ObjectID{header.ID},
		BodyRefs:   []primitive.ObjectID{body.ID},
	}

	// Insert components into their respective collections
	_, err = s.HeaderStorage.InsertOne(context.TODO(), header)
	if err != nil {
		return PostData{}, err
	}

	_, err = s.BodyStorage.InsertOne(context.TODO(), body)
	if err != nil {
		return PostData{}, err
	}

	_, err = s.ScaffoldStorage.InsertOne(context.TODO(), scaffold)
	if err != nil {
		return PostData{}, err
	}

	// Update post data before returning it
	result = post
	result.IDHex = scaffold.ID.String()
	result.CreatedAt = scaffold.CreatedAt
	return result, nil
}
