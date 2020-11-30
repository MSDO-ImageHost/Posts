package database

func AddPost(post Scaffold) (id string, err error) {
	id, err = storage.Add(post)
	if err != nil {
		// catch timeout errors
	}

	return id, err
}

func (s *mongoStorage) Add(post Scaffold) (id string, err error) {
	return id, nil
}
