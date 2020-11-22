package database

/*
// Find a post in the database with id
func (storage *MongoStorage) FindOne(postID string) (ReadPost, error) {

	return fetchedPost, nil
}


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
