package database

func DeleteManyPosts(postIdHexes []string) (results []PostData, err error) {
	results, err = storage.DeleteMany(postIdHexes)
	if err != nil {
		// TODO: catch timeout errors
	}
	return results, err
}

func (s *mongoStorage) DeleteMany(postIdHexes []string) (results []PostData, err error) {

	return results, nil
}
