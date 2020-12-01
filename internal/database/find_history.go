package database

func FindHistoryPost(postIdHex string) (result PostDataHistory, err error) {
	result, err = storage.FindHistory(postIdHex)
	if err != nil {
		// TODO: catch timeout errors
	}
	return result, err
}

func (s *mongoStorage) FindHistory(postIdHex string) (result PostDataHistory, err error) {

	return result, nil
}
