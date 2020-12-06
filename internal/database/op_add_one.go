package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Public module handler
func AddOnePost(post PostData) (result PostData, err error) {
	if err := AssertClientInstance(); err != nil {
		return result, err
	}
	return storage.AddOne(post)
}

func (s *mongoStorage) AddOne(post PostData) (result PostData, err error) {

	now := time.Now()

	// Construct post content components
	header := mongoContent{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Data:      post.Header.Data,
		CreatedAt: &now,
	}

	body := mongoContent{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Data:      post.Body.Data,
		CreatedAt: &now,
	}

	scaffold := mongoScaffoldRefs{
		ID:         primitive.NewObjectID(),
		AuthorID:   post.AuthorID,
		CreatedAt:  &now,
		HeaderRefs: []primitive.ObjectID{header.ID},
		BodyRefs:   []primitive.ObjectID{body.ID},
	}

	// Insert components into their respective collections
	_, err = s.HeaderStorage.InsertOne(context.TODO(), header)
	if err != nil {
		return result, err
	}

	_, err = s.BodyStorage.InsertOne(context.TODO(), body)
	if err != nil {
		return result, err
	}

	_, err = s.ScaffoldStorage.InsertOne(context.TODO(), scaffold)
	if err != nil {
		return result, err
	}

	// Update post data before returning it
	result = post
	result.IDHex = scaffold.ID.Hex()
	result.CreatedAt = scaffold.CreatedAt
	result.Header = PostContent{
		AuthorID:  header.AuthorID,
		Data:      header.Data,
		CreatedAt: header.CreatedAt,
	}
	result.Body = PostContent{
		AuthorID:  body.AuthorID,
		Data:      body.Data,
		CreatedAt: body.CreatedAt,
	}
	return result, nil
}
