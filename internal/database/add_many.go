package database

func AddManyPosts(posts []PostData) (results []PostData, err error) {
	results, err = storage.AddMany(posts)
	if err != nil {
		// TODO: catch timeout errors
	}

	return results, err
}

func (s *mongoStorage) AddMany(posts []PostData) (results []PostData, err error) {

	return results, nil
}
