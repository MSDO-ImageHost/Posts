package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Find a post in the database with id
func (storage *MongoStorage) FindOne(postID string) (PostScaffold, error) {

	var scaffold PostScaffold

	// Convert hex string into bson object id
	scaffoldID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return scaffold, err
	}

	scaffoldFilter := bson.M{"_id": scaffoldID}
	err = storage.ScaffoldStorage.FindOne(context.TODO(), scaffoldFilter).Decode(&scaffold)
	if err != nil {
		return scaffold, err
	}

	// Fetch latest post header scaffold.Header.([]primitive.ObjectID)
	var header Content
	headerFilter := bson.D{
		{"_id", bson.M{"$in": scaffold.Header}},
		//{"$orderby", bson.M{"created_at": -1}}, // TODO: sort by latest
	}
	err = storage.HeaderStorage.FindOne(context.TODO(), headerFilter).Decode(&header)
	if err != nil {
		return scaffold, err
	}

	// Fetch latest post body
	var body Content
	bodyFilter := bson.D{
		{"_id", bson.M{"$in": scaffold.Body}},
		//{"$orderby", bson.M{"created_at": -1}}, // TODO: sort by latest
	}

	err = storage.BodyStorage.FindOne(context.TODO(), bodyFilter).Decode(&body)
	if err != nil {
		return scaffold, err
	}

	// Compose fetched data
	scaffold.HeaderContent = header.Data
	scaffold.BodyContent = body.Data

	return scaffold, nil
}

func (storage *MongoStorage) FindOneHistory(postIDHex string) (PostScaffold, error) {
	return PostScaffold{}, nil
}

func (storage *MongoStorage) FindUserPosts(userId string) ([]PostScaffold, error) {
	return []PostScaffold{}, nil
}

/*
// Find a post in the database with id
func (storage *MongoStorage) FindMany(postID string) (fetchedPost models.Post, err error) {
	// Fetch post scaffold
	scaffoldID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return models.Post{}, err
	}
	var scaffold models.Post
	scaffoldFilter := bson.M{"_id": scaffoldID}
	err = coll.Handle.FindOne(context.TODO(), scaffoldFilter).Decode(&scaffold)
	if err != nil {
		return models.Post{}, err
	}

	// Fetch latest post header scaffold.Header.([]primitive.ObjectID)
	var header models.Content
	headerFilter := bson.D{
		{"_id", bson.M{"$in": scaffold.Header.(primitive.A)}},
		//{"$orderby", bson.M{"created_at": -1}}, // TODO: sort by latest
	}
	err = coll.Handle.FindOne(context.TODO(), headerFilter).Decode(&header)
	if err != nil {
		return models.Post{}, err
	}

	// Fetch latest post body
	var body models.Content
	bodyFilter := bson.D{
		{"_id", bson.M{"$in": scaffold.Body.(primitive.A)}},
		//{"$orderby", bson.M{"created_at": -1}}, // TODO: sort by latest
	}

	err = coll.Handle.FindOne(context.TODO(), bodyFilter).Decode(&body)
	if err != nil {
		return models.Post{}, err
	}

	// Compose fetched data
	fetchedPost = models.Post{
		ID:        scaffold.ID.(primitive.ObjectID).Hex(),
		AuthorID:  scaffold.AuthorID,
		CreatedAt: scaffold.CreatedAt,
		Header:    header.Data,
		Body:      body.Data,
	}

	return fetchedPost, nil
}
*/
