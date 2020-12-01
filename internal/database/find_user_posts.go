package database

func FindUserPosts(author string) (results []PostData, err error) {
	results, err = storage.FindUserPosts(author)
	if err != nil {
		// TODO: catch timeout errors
	}
	return results, err
}

func (s *mongoStorage) FindUserPosts(author string) (results []PostData, err error) {

	return results, nil
}
