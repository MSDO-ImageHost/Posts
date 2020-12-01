package database

func FindOnePost(postIdHex string) (result PostData, err error) {
	result, err = storage.FindOne(postIdHex)
	if err != nil {
		// TODO: catch timeout errors
	}
	return result, err
}

func (s *mongoStorage) FindOne(postIdHex string) (result PostData, err error) {

	return result, nil
}
