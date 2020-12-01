package database

func UpdateOnePost(post PostData) (result PostData, err error) {
	result, err = storage.UpdateOne(post)
	if err != nil {
		// TODO: catch timeout errors
	}
	return result, err
}

func (s *mongoStorage) UpdateOne(post PostData) (result PostData, err error) {

	return result, nil
}
