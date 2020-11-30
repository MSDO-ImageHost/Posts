package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Find a post in the database with id
func (storage *MongoStorage) FindOne(postID string) (scaffold PostScaffold, err error) {

	log.Println(postID)

	// Convert hex string into bson object id
	scaffoldID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return scaffold, err
	}

	findPostAggregation := bson.D{
		bson.M{ ''}
		bson.M{'$match': bson.M{'$expr': bson.M{'$in': ['$_id', '$$header_ids']}}},
	}

	_, err := storage.ScaffoldStorage.Aggregate(context.TODO(), findPostAggregation)
	if err != nil {
		return scaffold, err
	}

	return scaffold, nil
}

func (storage *MongoStorage) FindOneHistory(postIDHex string) (PostScaffold, error) {
	return PostScaffold{}, nil
}

func (storage *MongoStorage) FindUserPosts(userId string) (posts []PostScaffold, err error) {

	//scaffoldFilter := bson.M{"author_id": userId}
	//
	//cur, err := storage.ScaffoldStorage.Find(context.TODO(), scaffoldFilter)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for cur.Next(context.TODO()) {
	//	var result bson.M
	//	err := cur.Decode(&result)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	posts = append(posts, PostScaffold{
	//		ID: result.,
	//	})
	//}
	//if err := cur.Err(); err != nil {
	//	log.Fatal(err)
	//}

	return posts, nil
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
