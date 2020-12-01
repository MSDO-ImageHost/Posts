package database

func DeleteOnePost(postIdHex string) (result PostData, err error) {
	result, err = storage.DeleteOne(postIdHex)
	if err != nil {
		// TODO: catch timeout errors
	}
	return result, err
}

func (s *mongoStorage) DeleteOne(postIdHex string) (result PostData, err error) {

	return result, nil
}
