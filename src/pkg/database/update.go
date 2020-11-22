package database

/*
// Update
func (coll *ScaffoldStorage) Update(postID string, post PostScaffold) (updatedPostID string, err error) {
	now := time.Now()

	// Construct post components
	header := Content{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Data:      fmt.Sprintf("%v", post.Header),
		CreatedAt: now,
	}

	body := Content{
		ID:        primitive.NewObjectID(),
		AuthorID:  post.AuthorID,
		Data:      fmt.Sprintf("%v", post.Body),
		CreatedAt: now,
	}

	id, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return "", err
	}

	// Update scaffold
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"header_ids": header.ID, "body_ids": body.ID}}
	_, err = coll.Handle.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "", err
	}

	// Insert components into their respective collections
	_, err = Headers.Add(header)
	if err != nil {
		return "", err
	}

	_, err = Bodies.Add(body)
	if err != nil {
		return "", err
	}
	return updatedPostID, nil
}
*/
