package database

func FindManyPosts(postIdHexes []string) (results []PostData, err error) {
	results, err = storage.FindMany(postIdHexes)
	if err != nil {
		// TODO: catch timeout errors
	}
	return results, err
}

func (s *mongoStorage) FindMany(postIdHexes []string) (results []PostData, err error) {

	return results, nil
}
