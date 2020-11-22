package database

/*
// Delete
func (coll *ScaffoldStorage) Delete(postID string) (deletedPostID string, err error) {
	// Fetch and delete the post scaffold
	scaffoldID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return "", err
	}
	var scaffold models.Post
	scaffoldFilter := bson.M{"_id": scaffoldID}
	err = coll.Handle.FindOneAndDelete(context.TODO(), scaffoldFilter).Decode(&scaffold)
	if err != nil {
		return "", err
	}

	// Delete all belonging post headers
	headerFilter := bson.M{"_id": bson.M{"$in": scaffold.Header.(primitive.A)}}
	_, err = coll.Handle.DeleteMany(context.TODO(), headerFilter)
	if err != nil {
		return "", err
	}

	// Delete all belonging post bodies
	bodyFilter := bson.M{"_id": bson.M{"$in": scaffold.Body.(primitive.A)}}
	_, err = coll.Handle.DeleteMany(context.TODO(), bodyFilter)
	if err != nil {
		return "", err
	}

	return deletedPostID, nil
}
*/
